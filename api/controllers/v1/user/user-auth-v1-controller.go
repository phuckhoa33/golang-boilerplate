package user_auth_v1_controller

import (
	"fmt"
	user_auth_requests "golang-boilerplate/domain/requests/user/auth"
	wrapper_responses "golang-boilerplate/domain/responses"
	user_auth_responses "golang-boilerplate/domain/responses/user/auth"
	"golang-boilerplate/models"
	respositories "golang-boilerplate/respositories/postgresql"
	"golang-boilerplate/server"
	token_service "golang-boilerplate/services/user/token"
	"net/http"
	"strings"

	jwtGo "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthV1Controller struct {
	server          *server.Server
	userRespository *respositories.UserRepository
	tokenService    *token_service.UserTokenService
}

func NewUserAuthV1Controller(server *server.Server) *UserAuthV1Controller {
	return &UserAuthV1Controller{
		server:          server,
		userRespository: respositories.NewUserRepository(server.DB),
		tokenService:    token_service.NewTokenService(server.Config),
	}
}

// Login godoc
//
//	@Summary		Authenticate a user
//	@Description	Perform user login
//	@ID				user-login
//	@Tags			user.auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		user_auth_requests.LoginRequest	true	"User's credentials"
//	@Success		200		{object}	user_auth_responses.LoginResponse
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/user/login [post]
func (controller *UserAuthV1Controller) Login(context *gin.Context) {
	// Get request
	loginRequest := new(user_auth_requests.LoginRequest)

	if err := context.Bind(&loginRequest); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Check field is empty
	if err := loginRequest.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Required fields is empty or invalid")
		return
	}

	// Check user is existed
	user := models.User{}

	// Check user is existed
	controller.userRespository.GetUserByEmail(&user, loginRequest.Email)

	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil) {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Invalid credentials")
		return
	}

	// Initilize token service
	// tokenService := services.NewTokenService(controller.server.Config)

	// Create access token
	accessToken, exp, err := controller.tokenService.CreateAccessToken(&user)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Create refresh token
	refreshToken, err := controller.tokenService.CreateRefreshToken(&user)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}
	//  return tokens
	res := user_auth_responses.NewLoginResponse(accessToken, refreshToken, exp)
	wrapper_responses.Response(context, http.StatusOK, res)
}

// Register godoc
//
//	@Summary		Authenticate a user
//	@Description	Perform user login
//	@ID				user-register
//	@Tags			user.auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		user_auth_requests.RegisterRequest	true	"User's credentials"
//	@Success		200		{string} 	Register successufully
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/user/register [post]
func (controller *UserAuthV1Controller) Register(context *gin.Context) {
	// Get request information
	registerRequest := new(user_auth_requests.RegisterRequest)

	if err := context.Bind(&registerRequest); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Validate field for request
	if err := registerRequest.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check user is existed
	user := models.User{}
	controller.userRespository.GetUserByEmail(&user, registerRequest.Email)
	if user.ID != 0 {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, "User is existed")
		return
	}

	// Hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	// Create new user
	newUser := models.User{
		Username: registerRequest.Username,
		Email:    strings.ToLower(registerRequest.Email),
		Password: string(hashPassword),
		Fullname: registerRequest.Fullname,
		Gender:   registerRequest.Gender,
		RoleId:   1,
	}
	controller.server.DB.Create(&newUser)

	// Return response
	wrapper_responses.Response(context, http.StatusOK, "Register successfully")
}

// RefreshToken godoc
//
//	@Summary		Refresh access token
//	@Description	Perform refresh access token
//	@ID				user-refresh-token
//	@Tags			user.auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		user_auth_requests.RefreshRequest	true	"Refresh token"
//	@Success		200		{object}	user_auth_responses.LoginResponse
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/refresh-token [post]
func (controller *UserAuthV1Controller) RefeshToken(context *gin.Context) {
	// Initilizate request
	refreshTokenRequest := new(user_auth_requests.RefreshRequest)
	if err := context.Bind(&refreshTokenRequest); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Parse token
	token, err := jwtGo.Parse(refreshTokenRequest.Token, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED_SIGNING_METHOD: %v", token.Header["alg"])
		}
		return []byte(controller.server.Config.Auth.RefreshTokenSecret), nil
	})
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	claims, ok := token.Claims.(jwtGo.MapClaims)
	if !ok && !token.Valid {
		wrapper_responses.ErrorResponse(context, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Check user
	user := new(models.User)
	controller.server.DB.First(&user, int(claims["id"].(float64)))

	if user.ID == 0 {
		wrapper_responses.ErrorResponse(context, http.StatusUnauthorized, "User not found")
		return
	}

	// Create access token
	accessToken, exp, err := controller.tokenService.CreateAccessToken(user)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Create refresh token
	refreshToken, err := controller.tokenService.CreateRefreshToken(user)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}
	//  return tokens
	res := user_auth_responses.NewLoginResponse(accessToken, refreshToken, exp)
	wrapper_responses.Response(context, http.StatusOK, res)
}

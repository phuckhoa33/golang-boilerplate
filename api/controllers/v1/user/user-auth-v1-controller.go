package user_auth_v1_controller

import (
	user_auth_requests "golang-boilerplate/domain/requests/user/auth"
	wrapper_responses "golang-boilerplate/domain/responses"
	responses_user_auth "golang-boilerplate/domain/responses/user/auth"
	"golang-boilerplate/models"
	respositories "golang-boilerplate/respositories/postgresql"
	"golang-boilerplate/server"
	services "golang-boilerplate/services/user/token"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthV1Controller struct {
	server *server.Server
}

func NewUserAuthV1Controller(server *server.Server) *UserAuthV1Controller {
	return &UserAuthV1Controller{server: server}
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

	if err := context.Bind(loginRequest); err != nil {
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
	userRepository := respositories.NewUserRepository(controller.server.DB)

	// Check user is existed
	userRepository.GetUserByEmail(&user, loginRequest.Email)

	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil) {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Invalid credentials")
		return
	}

	// Initilize token service
	tokenService := services.NewTokenService(controller.server.Config)

	// Create access token
	accessToken, exp, err := tokenService.CreateAccessToken(&user)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Create refresh token
	refreshToken, err := tokenService.CreateRefreshToken(&user)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}
	//  return tokens
	res := responses_user_auth.NewLoginResponse(accessToken, refreshToken, exp)
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
	if err := registerRequest.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check user is existed
	user := models.User{}
	userRespository := respositories.NewUserRepository(controller.server.DB)
	userRespository.GetUserByEmail(&user, user.Email)
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
		Username:    registerRequest.Username,
		Email:       strings.ToLower(registerRequest.Email),
		Password:    string(hashPassword),
		PhoneNumber: "",
		Fullname:    registerRequest.Fullname,
		Address:     "",
		Gender:      registerRequest.Gender,
		DateOfBirth: time.Now(),
		Avatar:      "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	controller.server.DB.Create(newUser)

	// Return response
	wrapper_responses.Response(context, http.StatusOK, "Register successfully")
}

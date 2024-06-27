package user_auth_v1_controller

import (
	"fmt"
	user_auth_requests "golang-boilerplate/domain/requests/user/auth"
	wrapper_responses "golang-boilerplate/domain/responses"
	user_auth_responses "golang-boilerplate/domain/responses/user/auth"
	"golang-boilerplate/models"
	respositories "golang-boilerplate/respositories/postgresql"
	"golang-boilerplate/server"
	mail_service "golang-boilerplate/services/mail"
	random_creation_service "golang-boilerplate/services/shared"
	token_service "golang-boilerplate/services/user/token"
	"net/http"
	"strings"
	"time"

	jwtGo "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthV1Controller struct {
	server                *server.Server
	userRespository       *respositories.UserRepository
	tokenService          *token_service.UserTokenService
	mailService           *mail_service.MailService
	randomCreationService *random_creation_service.RandomCreationService
}

func NewUserAuthV1Controller(server *server.Server) *UserAuthV1Controller {
	return &UserAuthV1Controller{
		server:                server,
		userRespository:       respositories.NewUserRepository(server.DB),
		tokenService:          token_service.NewTokenService(server.Config),
		randomCreationService: random_creation_service.NewRandomCreationService(),
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

// ForgotPassword godoc
//
//	@Summary	 	Forgot password
//	@Description	Perform forgot password
//	@ID				user-forgot-password
//	@Tags			user.auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		user_auth_requests.ForgotPasswordRequest	true	"User email"
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/forgot-password [post]
func (controller *UserAuthV1Controller) ForgotPassword(context *gin.Context) {
	// Initilizate request
	forgotPasswordRequest := new(user_auth_requests.ForgotPasswordRequest)
	if err := context.Bind(&forgotPasswordRequest); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Validate request
	if err := forgotPasswordRequest.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check user
	user := new(models.User)
	controller.server.DB.Where("email = ?", forgotPasswordRequest.Email).First(&user)
	if user.ID == 0 {
		wrapper_responses.ErrorResponse(context, http.StatusNotFound, "User not found")
		return
	}

	// Create random otp for user
	otp := controller.randomCreationService.GenerateOTP(6)

	// Update verfifyAccountOtp in database of user
	controller.server.DB.Update("VerifyAccountOtp", otp).First(&user)

	// Create token have expired time
	token, err := controller.tokenService.CreateFogotPasswordToken(forgotPasswordRequest.Email, "Forgot Password", time.Minute*3)
	// Check generate token error
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Config to, subject, templateFile string, data interface{} for send email
	resetLink := fmt.Sprintf("%s:%s/reset-password?token=%s", controller.server.Config.App.AppHost, controller.server.Config.App.AppPort, token)
	subject := "Forgot password"
	templateFile := "forgot-password.html"
	data := map[string]interface{}{
		"username":  user.Username,
		"resetLink": resetLink,
	}

	// Send email
	err = controller.mailService.SendEmail(forgotPasswordRequest.Email, subject, templateFile, data)

	// Check error
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Return response
	wrapper_responses.Response(context, http.StatusOK, "Send Email Successfully")
}

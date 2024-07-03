package user_v1_controller

import (
	"fmt"
	"github.com/google/uuid"
	user_requests "golang-boilerplate/domain/requests/user"
	wrapper_responses "golang-boilerplate/domain/responses"
	user_responses "golang-boilerplate/domain/responses/user"
	"golang-boilerplate/models"
	respositories "golang-boilerplate/respositories/postgresql"
	"golang-boilerplate/server"
	mail_service "golang-boilerplate/services/mail"
	random_creation_service "golang-boilerplate/services/shared"
	token_service "golang-boilerplate/services/token"
	"net/http"
	"strings"
	"time"

	jwtGo "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthV1Controller struct {
	server                *server.Server
	userRepository        *respositories.UserRepository
	tokenService          *token_service.TokenService
	mailService           *mail_service.MailService
	randomCreationService *random_creation_service.RandomCreationService
}

func NewUserAuthV1Controller(server *server.Server) *UserAuthV1Controller {
	return &UserAuthV1Controller{
		server:                server,
		userRepository:        respositories.NewUserRepository(server.DB),
		tokenService:          token_service.NewTokenService(server.Config),
		mailService:           mail_service.NewMailService(server.Config),
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
//	@Param			params	body		user_requests.LoginRequest	true	"User's credentials"
//	@Success		200		{object}	user_responses.LoginResponse
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/user/login [post]
func (controller *UserAuthV1Controller) Login(context *gin.Context) {
	// Get request
	request := new(user_requests.LoginRequest)

	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Check field is empty
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Required fields is empty or invalid")
		return
	}

	// Check user is existed
	user := models.User{}

	// Check user is existed
	controller.userRepository.GetUserByEmail(&user, request.Email)

	if user.ID == uuid.Nil || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil) {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Invalid credentials")
		return
	}

	// Initialize token service
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
	res := user_responses.NewLoginResponse(accessToken, refreshToken, exp)
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
//	@Param			params	body		user_requests.RegisterRequest	true	"User's credentials"
//	@Success		200		{string}	string	"Register successfully"
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/user/register [post]
func (controller *UserAuthV1Controller) Register(context *gin.Context) {
	// Get request information
	request := new(user_requests.RegisterRequest)

	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Validate field for request
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check user is existed
	user := models.User{}
	controller.userRepository.GetUserByEmail(&user, request.Email)
	if user.ID != uuid.Nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, "User is existed")
		return
	}

	// Hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	// Create new user
	newUser := models.User{
		Username: request.Username,
		Email:    strings.ToLower(request.Email),
		Password: string(hashPassword),
		FullName: request.FullName,
		Gender:   request.Gender,
	}
	controller.userRepository.Create(&newUser)

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
//	@Param			params	body		user_requests.RefreshRequest	true	"Refresh token"
//	@Success		200		{object}	user_responses.LoginResponse
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/user/refresh-token [post]
func (controller *UserAuthV1Controller) RefreshToken(context *gin.Context) {
	// Initialize request
	request := new(user_requests.RefreshRequest)
	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Parse token
	token, err := jwtGo.Parse(request.Token, func(token *jwtGo.Token) (interface{}, error) {
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
	user := models.User{}
	controller.userRepository.GetUserById(&user, int(claims["id"].(float64)))

	if user.ID == uuid.Nil {
		wrapper_responses.ErrorResponse(context, http.StatusUnauthorized, "User not found")
		return
	}

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
	res := user_responses.NewLoginResponse(accessToken, refreshToken, exp)
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
//	@Param			params	body		user_requests.ForgotPasswordRequest	true	"User email"
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/user/forgot-password [post]
func (controller *UserAuthV1Controller) ForgotPassword(context *gin.Context) {
	// Initialize request
	request := new(user_requests.ForgotPasswordRequest)
	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Validate request
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check user
	user := models.User{}
	controller.userRepository.GetUserByEmail(&user, request.Email)
	if user.ID == uuid.Nil {
		wrapper_responses.ErrorResponse(context, http.StatusNotFound, "User not found")
		return
	}

	// Create random otp for user
	otp := controller.randomCreationService.GenerateOTP(6)

	// Update verifyAccountOtp in database of user
	controller.userRepository.UpdateSingleProperty(&user, "verifyAccountOtp", otp)

	// Create token have expired time
	token, err := controller.tokenService.CreateForgotPasswordToken(request.Email, "Forgot Password", time.Minute*3)
	// Check generate token error
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Config to, subject, templateFile string, data interface{} for send email
	resetLink := fmt.Sprintf("%s:%s/reset-password?token=%s", controller.server.Config.App.AppHost, controller.server.Config.App.AppPort, token)
	subject := "Forgot password"
	templateFile := "templates/mail/forgot-password.html"
	data := map[string]interface{}{
		"Username":  user.Username,
		"ResetLink": resetLink,
		"AppName":   controller.server.Config.App.AppName,
	}

	// Send email
	err = controller.mailService.SendEmail(request.Email, subject, templateFile, data)

	// Check error
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Send email error")
		return
	}

	// Return response
	wrapper_responses.Response(context, http.StatusOK, "Send Email Successfully")
}

// CheckValidForgotPasswordLink godoc
//
//	@Summary	 	Check forgot password link
//	@Description	Perform Check forgot password link
//	@ID				check-valid-forgot-password-link
//	@Tags			user.auth
//	@Accept			json
//	@Produce		json
//	@Param token path string true "The forgot password token sent to the user's email."
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/user/check-valid-forgot-link/{token} [get]
func (controller *UserAuthV1Controller) CheckValidForgotPasswordLink(context *gin.Context) {
	param := context.Param("token")

	// Parse token
	token, err := jwtGo.Parse(param, func(token *jwtGo.Token) (interface{}, error) {
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
	user := models.User{}
	controller.userRepository.GetUserById(&user, claims["iss"])
	if user.ID == uuid.Nil {
		wrapper_responses.ErrorResponse(context, http.StatusNotFound, "User not found")
		return
	}

	// Return response
	wrapper_responses.Response(context, http.StatusOK, true)
}

// ForgotPasswordToResetPassword godoc
//
//	@Summary	 	Forgot password to reset password
//	@Description	Perform forgot password to reset password
//	@ID				user-forgot-password-to-reset-password
//	@Tags			user.auth
//	@Accept			json
//	@Produce		json
//	@Param token path string true "The forgot password token sent to the user's email."
//	@Param			params	body		user_requests.ForgotPasswordToResetPasswordRequest	true	"New password of user"
//	@Failure		401		{object}	wrapper_responses.Error
//	@Router			/user/forgot-password/{token} [post]
func (controller *UserAuthV1Controller) ForgotPasswordToResetPassword(context *gin.Context) {
	request := new(user_requests.ForgotPasswordToResetPasswordRequest)
	param := context.Param("token")

	// Parse token
	token, err := jwtGo.Parse(param, func(token *jwtGo.Token) (interface{}, error) {
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
	user := models.User{}
	controller.userRepository.GetUserById(&user, claims["iss"])
	if user.ID == uuid.Nil {
		wrapper_responses.ErrorResponse(context, http.StatusNotFound, "User not found")
		return
	}

	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Validate request
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check two password equal
	if request.ConfirmNewPassword != request.NewPassword {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Confirm new password must be equal new password")
		return
	}

	// Reset password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	controller.userRepository.UpdateSingleProperty(&user, "password", string(hashPassword))

	// Return response
	wrapper_responses.Response(context, http.StatusOK, "Reset password successfully")
}

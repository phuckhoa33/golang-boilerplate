package user_v1_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	user_requests "golang-boilerplate/domain/requests/user"
	wrapper_responses "golang-boilerplate/domain/responses"
	user_responses "golang-boilerplate/domain/responses/user"
	"golang-boilerplate/models"
	respositories "golang-boilerplate/respositories/postgresql"
	"golang-boilerplate/server"
	minio_service "golang-boilerplate/services/minio"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserMeV1Controller struct {
	server             *server.Server
	userRepository     *respositories.UserRepository
	minioClientService *minio_service.MinioClientService
}

func NewUserMeV1Controller(server *server.Server) *UserMeV1Controller {
	return &UserMeV1Controller{
		server:             server,
		userRepository:     respositories.NewUserRepository(server.DB),
		minioClientService: minio_service.NewMinioClientService(server.Config),
	}
}

// GetUserProfile @Summary Get user profile
// @Description Get user profile
// @ID get-user-profile
// @Tags user.me
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} user_responses.ViewUserProfileResponse
// @Router /user/me [get]
func (controller *UserMeV1Controller) GetUserProfile(context *gin.Context) {
	// Get user id from context
	userId, _ := context.Get("userId")

	// Get user profile
	user := models.User{}
	controller.userRepository.GetUserById(&user, userId)

	// Check user is existed
	if user.ID == uuid.Nil {
		wrapper_responses.ErrorResponse(context, http.StatusNotFound, "NOT_FOUND_USER")
		return
	}

	// Mapping response
	res := user_responses.NewViewUserProfileResponse(&user)

	wrapper_responses.Response(context, http.StatusOK, res)
}

// GetPreSignedPutURL  @Summary Get pre signed put url
// @Description Get pre signed put url of minio cloud(cloud for images and files)
// @ID get-pre-signed-url
// @Tags user.me
// @Accept json
// @Produce json
// @Param objectName path string true "Object name"
// @Security ApiKeyAuth
// @Success 200 {object} user_responses.GetUserPutPreSignedPutURLResponse
// @Router /user/pre-signed-put-url/{objectName} [get]
func (controller *UserMeV1Controller) GetPreSignedPutURL(context *gin.Context) {
	param := context.Param("objectName")
	newURL, err := controller.minioClientService.GetPutPreSignedURL(context, param)
	if err != nil {
		panic(err)
	}
	res := user_responses.NewGetUserPutPreSignedPutURLResponse(newURL)
	wrapper_responses.Response(context, http.StatusOK, res)
}

// UpdateUserInfo godoc
//
//	@Summary		Update user info
//	@Description	Update user info with basic user information
//	@ID				user-update-user-info
//	@Tags			user.me
//	@Accept			json
//	@Produce		json
//	@Param			params	body		user_requests.UpdateUserInfoRequest	true	"Update User Info"
//	@Success		200
//	@Security ApiKeyAuth
//	@Failure		400		{object}	wrapper_responses.Error
//	@Router			/user [put]
func (controller *UserMeV1Controller) UpdateUserInfo(context *gin.Context) {
	// Get user id from context
	userId, _ := context.Get("userId")

	// Get body request
	request := new(user_requests.UpdateUserInfoRequest)
	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Get user and check is existed
	user := models.User{}
	controller.userRepository.GetUserById(&user, userId)
	if user.ID == uuid.Nil {
		wrapper_responses.ErrorResponse(context, http.StatusNotFound, "NOT_FOUND_USER")
		return
	}

	// Update user information
	// TODO: Lack of some fields for updating
	user.Gender = request.Gender
	user.Address = request.Address
	user.FullName = request.FullName
	user.PhoneNumber = request.PhoneNumber
	user.Username = request.Username

	controller.userRepository.UpdateUser(&user)
}

// ChangePassword godoc
//
//	@Summary		Change password
//	@Description	Change password for user with old password and new password
//	@ID				user-change-password
//	@Tags			user.me
//	@Accept			json
//	@Produce		json
//	@Param			params	body		user_requests.ChangePasswordRequest	true	"Change password"
//	@Success		200
//	@Security ApiKeyAuth
//	@Failure		400		{object}	wrapper_responses.Error
//	@Router			/user/change-password [patch]
func (controller *UserMeV1Controller) ChangePassword(context *gin.Context) {
	// Get user id from context
	userId, _ := context.Get("userId")
	// get body
	request := new(user_requests.ChangePasswordRequest)

	// Bind value
	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadGateway, err.Error())
		return
	}

	// Check field is empty
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "INVALID_INPUT_FORMAT")
		return
	}

	// Check user
	user := models.User{}
	controller.userRepository.GetUserById(&user, userId)
	// Check user
	if user.ID == uuid.Nil {
		wrapper_responses.ErrorResponse(context, http.StatusNotFound, "NOT_FOUND_USER")
		return
	}

	// Check old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword)); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusNotFound, "PASSWORD_IS_INCORRECT")
		return
	}

	// Update new password
	hashNewPassword, _ := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	controller.userRepository.UpdateSingleProperty(&user, "password", string(hashNewPassword))
}

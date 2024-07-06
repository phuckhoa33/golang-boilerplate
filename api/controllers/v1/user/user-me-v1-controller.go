package user_v1_controller

import (
	"encoding/json"
	"github.com/google/uuid"
	wrapper_responses "golang-boilerplate/domain/responses"
	user_responses "golang-boilerplate/domain/responses/user"
	"golang-boilerplate/models"
	respositories "golang-boilerplate/respositories/postgresql"
	"golang-boilerplate/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserMeV1Controller struct {
	server         *server.Server
	userRepository *respositories.UserRepository
}

func NewUserMeV1Controller(server *server.Server) *UserMeV1Controller {
	return &UserMeV1Controller{
		server:         server,
		userRepository: respositories.NewUserRepository(server.DB),
	}
}

// GetUserProfile @Summary Get user profile
// @Description Get user profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} user_responses.UserProfileResponse
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
	var response user_responses.UserProfileResponse
	userBytes, err := json.Marshal(user)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.Unmarshal(userBytes, &response.Data); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	wrapper_responses.Response(context, http.StatusOK, response)
}

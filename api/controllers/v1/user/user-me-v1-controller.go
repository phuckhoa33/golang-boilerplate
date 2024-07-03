package user_v1_controller

import (
	"encoding/json"
	wrapper_responses "golang-boilerplate/domain/responses"
	user_responses "golang-boilerplate/domain/responses/user"
	"golang-boilerplate/models"
	"golang-boilerplate/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserMeV1Controller struct {
	server *server.Server
}

func NewUserMeV1Controller(server *server.Server) *UserMeV1Controller {
	return &UserMeV1Controller{
		server: server,
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
	userId, existed := context.Get("userId")
	if !existed {
		wrapper_responses.ErrorResponse(context, http.StatusInternalServerError, "Internal server error")
		return
	}
	userIdStr := userId.(float64)

	// Get user profile
	user := models.User{}
	if err := controller.server.DB.Where("id = ?", userIdStr).First(&user).Error; err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusUnauthorized, err.Error())
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

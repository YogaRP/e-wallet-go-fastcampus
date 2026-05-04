package api

import (
	"ewallet-fastcampus/constants"
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/interfaces"
	"ewallet-fastcampus/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log  = helpers.Logger
		req  models.LoginRequest
		resp *models.LoginResponse
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	resp, err := api.LoginService.Login(c.Request.Context(), &req)

	if err != nil {
		log.Error("failed to login service: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}

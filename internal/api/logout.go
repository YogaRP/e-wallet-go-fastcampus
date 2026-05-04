package api

import (
	"ewallet-fastcampus/constants"
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token := c.Request.Header.Get("Authorization")
	err := api.LogoutService.Logout(c.Request.Context(), token)

	if err != nil {
		log.Error("failed to login service: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}

package api

import (
	"ewallet-fastcampus/constants"
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRerfreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	refreshToken := c.Request.Header.Get("Authorization")
	claim, ok := c.Get("token")

	if !ok {
		log.Error("failed to get claim in context")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("failed to get claim in context")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.RefreshTokenService.RefreshToken(c.Request.Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Error("failed to login service: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}

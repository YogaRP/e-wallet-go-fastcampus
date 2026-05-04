package interfaces

import (
	"context"
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/models"

	"github.com/gin-gonic/gin"
)

type IRerfreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (*models.RefreshTokenResponse, error)
}

type IRerfreshTokenHandler interface {
	RefreshToken(c *gin.Context)
}

package interfaces

import (
	"context"
	"ewallet-fastcampus/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (any, error)
}

type IRegisterHandler interface {
	RegisterHandler(c *gin.Context)
}

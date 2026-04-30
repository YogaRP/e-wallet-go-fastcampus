package interfaces

import (
	"context"
	"ewallet-fastcampus/internal/models"
)

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (any, error)
}

type IRegisterRepo interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}

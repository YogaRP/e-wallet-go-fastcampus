package services

import (
	"context"
	"ewallet-fastcampus/internal/interfaces"
)

type LogoutService struct {
	UserRepo interfaces.IUserRepo
}

func (r *LogoutService) Logout(ctx context.Context, token string) error {
	return r.UserRepo.DeleteUserSession(ctx, token)
}

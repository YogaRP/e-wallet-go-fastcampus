package services

import (
	"context"
	"ewallet-fastcampus/internal/interfaces"
	"ewallet-fastcampus/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo interfaces.IUserRepo
}

func (s *UserService) Register(ctx context.Context, request models.User) (any, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	request.Password = string(hashPassword)

	err = s.UserRepo.InsertNewUser(ctx, &request)

	if err != nil {
		return nil, err
	}

	resp := request
	resp.Password = ""
	return resp, nil
}

package services

import (
	"context"
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/interfaces"
	"ewallet-fastcampus/internal/models"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepo
}

func (r *LoginService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	var (
		response models.LoginResponse
		now      = time.Now()
	)

	userDetail, err := r.UserRepo.GetUserByUsername(ctx, req.Username)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return nil, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", now)

	if err != nil {
		return nil, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "refresh_token", now)

	if err != nil {
		return nil, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken[token]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken[refreshToken]),
	}

	err = r.UserRepo.InsertNewUserSession(ctx, userSession)

	if err != nil {
		return nil, errors.Wrap(err, "failed to insert new token")
	}

	response.UserID = userDetail.ID
	response.Username = userDetail.Username
	response.FullName = userDetail.FullName
	response.Email = userDetail.Email
	response.Token = token
	response.RefreshToken = refreshToken

	return &response, nil
}

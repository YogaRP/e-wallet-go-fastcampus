package services

import (
	"context"
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/interfaces"
	"ewallet-fastcampus/internal/models"
	"time"

	"github.com/pkg/errors"
)

type RefreshTokenService struct {
	UserRepo interfaces.IUserRepo
}

func (r *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (*models.RefreshTokenResponse, error) {
	resp := models.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, tokenClaim.UserID, tokenClaim.Username, tokenClaim.Fullname, "token", time.Now())

	if err != nil {
		return nil, errors.Wrap(err, "failed to generate new token")
	}

	err = r.UserRepo.UpdateTokenByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update new token")
	}

	resp.Token = token
	return &resp, nil
}

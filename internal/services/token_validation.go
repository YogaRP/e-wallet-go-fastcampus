package services

import (
	"context"
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/interfaces"

	"github.com/pkg/errors"
)

type TokenValidationService struct {
	UserRepo interfaces.IUserRepo
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var (
		claimToken *helpers.ClaimToken
	)

	claimToken, err := helpers.ValidateToken(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate token")
	}

	_, err = s.UserRepo.GetUserSessionByToken(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user session")
	}

	return claimToken, nil
}

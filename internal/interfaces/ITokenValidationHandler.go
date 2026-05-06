package interfaces

import (
	"context"
	"ewallet-fastcampus/cmd/proto/tokenvalidation"
	"ewallet-fastcampus/helpers"
)

type ITokenValidationHandler interface {
	ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}

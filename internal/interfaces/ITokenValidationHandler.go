package interfaces

import (
	"context"
	pb "ewallet-fastcampus/cmd/proto"
	"ewallet-fastcampus/helpers"
)

type ITokenValidationHandler interface {
	ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}

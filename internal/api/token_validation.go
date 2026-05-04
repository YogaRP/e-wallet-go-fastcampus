package api

import (
	"context"
	pb "ewallet-fastcampus/cmd/proto"
	"ewallet-fastcampus/constants"
	"ewallet-fastcampus/helpers"
	"ewallet-fastcampus/internal/interfaces"
	"fmt"
)

type TokenValidationHandler struct {
	TokenValidationService interfaces.ITokenValidationService
	pb.UnimplementedTokenValidationServer
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	var (
		token = req.Token
		log   = helpers.Logger
	)

	if token == "" {
		err := fmt.Errorf("token is empty")
		log.Error(err)
		return &pb.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		return &pb.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &pb.TokenResponse{
		Message: constants.SuccessMessage,
		Data: &pb.UserData{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.Fullname,
		},
	}, nil
}

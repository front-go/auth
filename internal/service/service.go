package service

import (
	"context"
	"github.com/front-go/auth/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	auth.UnimplementedAuthServiceServer
	dbR DbRepo
}

func NewService(dbR DbRepo) *Service {
	return &Service{dbR: dbR}
}

func (s *Service) Signup(ctx context.Context, in *auth.SignupIn) (*auth.SignupOut, error) {
	if in.Password != in.ConfirmPassword {
		return nil, status.Error(codes.FailedPrecondition, "password mismatch")
	}
	err := s.dbR.Insert(ctx, in.Username, in.Password)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "failed to insert user")
	}
	return &auth.SignupOut{
		Success: true,
	}, nil
}

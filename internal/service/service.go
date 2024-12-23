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

func (s *Service) Login(ctx context.Context, in *auth.LoginIn) (*auth.LoginOut, error) {
	storedPassword, err := s.dbR.GetPassword(ctx, in.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if storedPassword != in.Password {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	return &auth.LoginOut{
		Success: true,
	}, nil
}

func (s *Service) ChangePassword(ctx context.Context, in *auth.ChangePasswordIn) (*auth.ChangePasswordOut, error) {
	storedPassword, err := s.dbR.GetPassword(ctx, in.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if storedPassword != in.Password {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	response, err := s.dbR.UpdatePassword(ctx, in.Username, in.Password, in.NewPassword)

	return &auth.ChangePasswordOut{
		Response: response,
	}, err
}

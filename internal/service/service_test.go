package service

import (
	"context"
	"errors"
	"testing"

	"github.com/front-go/auth/internal/repository"
	"github.com/front-go/auth/pkg/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Signup(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDbRepo := NewMockDbRepo(ctrl)

	t.Run("ok", func(t *testing.T) {
		ctx := context.Background()
		in := &auth.SignupIn{
			Username:        "test",
			Password:        "123",
			ConfirmPassword: "123",
		}

		mockDbRepo.EXPECT().Insert(gomock.Any(), in.Username, in.Password).Return(nil)

		srv := NewService(mockDbRepo)
		_, err := srv.Signup(ctx, in)
		assert.NoError(t, err)
	})

	t.Run("password_not_match", func(t *testing.T) {
		ctx := context.Background()
		in := &auth.SignupIn{
			Password:        "123",
			ConfirmPassword: "321",
		}

		srv := NewService(mockDbRepo)
		_, err := srv.Signup(ctx, in)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password mismatch")
	})

	t.Run("fail call insert", func(t *testing.T) {
		ctx := context.Background()
		mockError := errors.New("mock error")
		in := &auth.SignupIn{
			Username:        "test",
			Password:        "123",
			ConfirmPassword: "123",
		}

		mockDbRepo.EXPECT().Insert(gomock.Any(), in.Username, in.Password).Return(mockError)

		srv := NewService(mockDbRepo)
		_, err := srv.Signup(ctx, in)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to insert user")
	})

	t.Run("fail_call_insert_err_already_exist", func(t *testing.T) {
		ctx := context.Background()
		mockError := repository.ErrAlreadyExist
		in := &auth.SignupIn{
			Username:        "test",
			Password:        "123",
			ConfirmPassword: "123",
		}

		mockDbRepo.EXPECT().Insert(gomock.Any(), in.Username, in.Password).Return(mockError)

		srv := NewService(mockDbRepo)
		_, err := srv.Signup(ctx, in)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Пользователь уже существует")
	})
}

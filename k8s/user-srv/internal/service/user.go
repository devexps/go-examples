package service

import (
	"context"

	v1 "github.com/devexps/go-examples/k8s/api/gen/go/user-srv/v1"
	
	"github.com/devexps/go-examples/k8s/user-srv/internal/biz"
)

type UserService interface {
	v1.UserServiceServer
}

type userService struct {
	v1.UnimplementedUserServiceServer

	uc biz.UserUseCase
}

// NewUserService new a user service.
func NewUserService(uc biz.UserUseCase) UserService {
	return &userService{uc: uc}
}

// VerifyPassword .
func (s *userService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.User, error) {
	user, err := s.uc.VerifyPassword(ctx, req.GetUserName(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return user, nil
}

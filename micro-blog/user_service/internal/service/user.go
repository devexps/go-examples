package service

import (
	"context"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/user_service/v1"
	"github.com/devexps/go-examples/micro-blog/user_service/internal/biz"
	"github.com/devexps/go-micro/v2/log"
)

// UserService interface
type UserService interface {
	v1.UserServiceServer
}

// userService is a user service.
type userService struct {
	v1.UnimplementedUserServiceServer

	log *log.Helper
	uc  biz.UserUseCase
}

// NewUserService new a greeter service.
func NewUserService(logger log.Logger, uc biz.UserUseCase) UserService {
	l := log.NewHelper(log.With(logger, "module", "user_service/service/user"))
	return &userService{
		log: l,
		uc:  uc,
	}
}

// VerifyPassword .
func (s *userService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.User, error) {
	user, err := s.uc.VerifyPassword(ctx, req.GetUserName(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser .
func (s *userService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	user, err := s.uc.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return user, nil
}

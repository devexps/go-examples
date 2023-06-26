package biz

import (
	"context"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/user_service/v1"
	"github.com/devexps/go-examples/micro-blog/user_service/internal/data"
	"github.com/devexps/go-micro/v2/log"
)

const (
	tempPassword = "******"
)

// UserUseCase is a User use case interface.
type UserUseCase interface {
	VerifyPassword(context.Context, string, string) (*v1.User, error)
	Get(context.Context, string) (*v1.User, error)
}

// userUseCase is a User use case.
type userUseCase struct {
	log  *log.Helper
	repo data.UserRepo
}

// NewUserUseCase new a User use case.
func NewUserUseCase(logger log.Logger, repo data.UserRepo) UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "user_service/usecase/user"))
	return &userUseCase{
		log:  l,
		repo: repo,
	}
}

// VerifyPassword .
func (u *userUseCase) VerifyPassword(ctx context.Context, userName string, password string) (*v1.User, error) {
	user, err := u.repo.VerifyPassword(ctx, userName, password)
	if user != nil {
		user.Password = tempPassword
	}
	return user, err
}

// Get .
func (u *userUseCase) Get(ctx context.Context, id string) (*v1.User, error) {
	user, err := u.repo.Get(ctx, id)
	if user != nil {
		user.Password = tempPassword
	}
	return user, err
}

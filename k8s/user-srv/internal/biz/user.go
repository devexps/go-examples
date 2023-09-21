package biz

import (
	"context"

	v1 "github.com/devexps/go-examples/k8s/api/gen/go/user-srv/v1"
	"github.com/devexps/go-examples/k8s/user-srv/internal/data"
	
	"github.com/devexps/go-micro/v2/log"
)

const (
	tempPassword = "******"
)

type UserUseCase interface {
	VerifyPassword(context.Context, string, string) (*v1.User, error)
}

type userUseCase struct {
	log  *log.Helper
	repo data.UserRepo
}

// NewUserUseCase new a User use case.
func NewUserUseCase(logger log.Logger, repo data.UserRepo) UserUseCase {
	l := log.NewHelper(log.With(logger, "module", "user-srv/usecase/user"))
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

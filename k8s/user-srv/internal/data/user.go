package data

import (
	"context"

	v1 "github.com/devexps/go-examples/k8s/api/gen/go/user-srv/v1"
	
	"github.com/devexps/go-micro/v2/log"
)

// UserRepo is a User repo interface.
type UserRepo interface {
	VerifyPassword(context.Context, string, string) (*v1.User, error)
}

// userRepo is a User repository
type userRepo struct {
	log  *log.Helper
	data *Data
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user-srv/data/user"))
	return &userRepo{
		data: data,
		log:  l,
	}
}

// VerifyPassword .
func (r *userRepo) VerifyPassword(ctx context.Context, userName string, password string) (*v1.User, error) {
	if userName != "thangn" || password != "123456" {
		return nil, v1.ErrorUserNotFound("User not found")
	}
	return &v1.User{
		Id:       "1",
		UserName: userName,
		NickName: userName,
		Password: password,
	}, nil
}

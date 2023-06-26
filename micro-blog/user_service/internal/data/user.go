package data

import (
	"context"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/user_service/v1"
	"github.com/devexps/go-examples/micro-blog/pkg/util/crypto"
	"github.com/devexps/go-examples/micro-blog/user_service/internal/data/ent"
	"github.com/devexps/go-examples/micro-blog/user_service/internal/data/ent/user"
	"github.com/devexps/go-micro/v2/log"
	"strconv"
)

// UserRepo is a User repo interface.
type UserRepo interface {
	VerifyPassword(context.Context, string, string) (*v1.User, error)
	Get(context.Context, string) (*v1.User, error)
}

// userRepo is a User repository
type userRepo struct {
	log  *log.Helper
	data *Data
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user_service/data/user"))
	return &userRepo{
		data: data,
		log:  l,
	}
}

// VerifyPassword .
func (r *userRepo) VerifyPassword(ctx context.Context, userName string, password string) (*v1.User, error) {
	u, err := r.data.db.User.
		Query().
		Where(user.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			r.log.Errorf("query user from db failed: %v", err.Error())
		}
		return nil, v1.ErrorUserNotFound("User not found")
	}
	bMatched := crypto.CheckPasswordHash(password, u.Password)
	if !bMatched {
		return nil, v1.ErrorUserNotFound("User not found")
	}
	return r.convertEntToProto(u), nil
}

// Get .
func (r *userRepo) Get(ctx context.Context, id string) (*v1.User, error) {
	res, err := r.data.db.User.Get(ctx, r.unit32FromString(id))
	if err != nil {
		if !ent.IsNotFound(err) {
			r.log.Errorf("get user from db failed: %v", err.Error())
		}
		return nil, v1.ErrorUserNotFound("User not found")
	}
	return r.convertEntToProto(res), err
}

func (r *userRepo) convertEntToProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}
	u := &v1.User{
		Id:       strconv.FormatUint(uint64(in.ID), 10),
		UserName: in.Username,
		NickName: in.Nickname,
		Password: in.Password,
	}
	return u
}

func (r *userRepo) unit32FromString(in string) uint32 {
	in_, _ := strconv.ParseUint(in, 10, 32)
	return uint32(in_)
}

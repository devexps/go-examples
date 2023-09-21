package data

import (
	"context"
	"time"

	jwtV4 "github.com/golang-jwt/jwt/v4"

	"github.com/devexps/go-examples/k8s/api/gen/go/common/authn"
	"github.com/devexps/go-examples/k8s/api/gen/go/common/conf"
	v1 "github.com/devexps/go-examples/k8s/api/gen/go/token-srv/v1"

	"github.com/devexps/go-micro/v2/log"
)

const (
	ClaimSubject  = "sub"
	ClaimIssuedAt = "iat"
	ClaimUserId   = "uid"
	ClaimPlatform = "platform"

	userTokenKeyPrefix = "ut_"
)

// TokenRepo is a Token repo interface.
type TokenRepo interface {
	GenerateToken(context.Context, *authn.Claims) (string, error)
}

// tokenRepo is a Token repository
type tokenRepo struct {
	log  *log.Helper
	data *Data
	key  []byte
}

// NewTokenRepo .
func NewTokenRepo(logger log.Logger, data *Data, cfg *conf.Bootstrap) TokenRepo {
	return &tokenRepo{
		log:  log.NewHelper(logger),
		data: data,
		key:  []byte(cfg.Auth.ApiKey),
	}
}

// GenerateToken .
func (r *tokenRepo) GenerateToken(ctx context.Context, info *authn.Claims) (string, error) {
	token := r.createJwtToken(ctx, info)
	if len(token) == 0 {
		return "", v1.ErrorCreateTokenFailed("create token failed")
	}
	if err := r.setToken(ctx, info); err != nil {
		r.log.Errorf("fail to generate token: %v", err)
		return "", v1.ErrorSaveTokenFailed("save token failed")
	}
	return token, nil
}

func (r *tokenRepo) createJwtToken(ctx context.Context, info *authn.Claims) string {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256,
		jwtV4.MapClaims{
			ClaimSubject:  info.GetUserName(),
			ClaimUserId:   info.GetUserId(),
			ClaimPlatform: info.GetPlatform(),
			ClaimIssuedAt: time.Now().Unix(),
		},
	)
	signedToken, err := claims.SignedString(r.key)
	if err != nil {
		r.log.Errorf("fail to create JWT token: %v", err)
		return ""
	}
	return signedToken
}

func (r *tokenRepo) setToken(ctx context.Context, info *authn.Claims) error {
	r.log.Infof("setToken: %v", info)
	// TODO: store token into redis
	return nil
}

package service

import (
	"context"

	v1 "github.com/devexps/go-examples/k8s/api/gen/go/token-srv/v1"
	"github.com/devexps/go-examples/k8s/token-srv/internal/biz"
	
	"github.com/devexps/go-micro/v2/log"
)

type TokenService interface {
	v1.TokenServiceServer
}

// TokenService is a token service.
type tokenService struct {
	v1.UnimplementedTokenServiceServer

	log *log.Helper
	uc  biz.TokenUseCase
}

// NewTokenService new a token service.
func NewTokenService(logger log.Logger, uc biz.TokenUseCase) TokenService {
	l := log.NewHelper(log.With(logger, "module", "token-srv/service/token"))
	return &tokenService{
		log: l,
		uc:  uc,
	}
}

// GenerateToken .
func (s *tokenService) GenerateToken(ctx context.Context, req *v1.GenerateTokenRequest) (*v1.GenerateTokenReply, error) {
	token, err := s.uc.GenerateToken(ctx, req.GetTokenInfo())
	if err != nil {
		s.log.Errorf("fail to generate token: %v", err)
		return nil, err
	}
	return &v1.GenerateTokenReply{
		Token: token,
	}, nil
}

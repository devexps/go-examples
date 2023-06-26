package service

import (
	"context"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/token_service/v1"
	"github.com/devexps/go-examples/micro-blog/token_service/internal/biz"
	"github.com/devexps/go-micro/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
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
	l := log.NewHelper(log.With(logger, "module", "token_service/service/token"))
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

// RemoveToken .
func (s *tokenService) RemoveToken(ctx context.Context, req *v1.RemoveTokenRequest) (*empty.Empty, error) {
	err := s.uc.RemoveToken(ctx, req.GetTokenInfo(), req.GetAll())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// ValidateToken .
func (s *tokenService) ValidateToken(ctx context.Context, req *v1.ValidateTokenRequest) (*v1.ValidateTokenReply, error) {
	retTokenInfo, err := s.uc.ValidateToken(ctx, req.GetToken())
	if err != nil {
		return nil, err
	}
	return &v1.ValidateTokenReply{
		TokenInfo: retTokenInfo,
	}, nil
}

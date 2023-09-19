package service

import (
	"context"
	v1 "github.com/devexps/go-examples/k8s/api/gen/go/admin-srv/v1"
	"github.com/devexps/go-examples/k8s/api/gen/go/common/authn"
	tokenV1 "github.com/devexps/go-examples/k8s/api/gen/go/token-srv/v1"
	userV1 "github.com/devexps/go-examples/k8s/api/gen/go/user-srv/v1"
	"github.com/devexps/go-micro/v2/log"
)

type AuthenticationService interface {
	v1.AuthenticationServiceServer
}

type authenticationService struct {
	v1.UnimplementedAuthenticationServiceServer

	log *log.Helper
	uc  userV1.UserServiceClient
	tc  tokenV1.TokenServiceClient
}

// NewAuthenticationService new an authentication service.
func NewAuthenticationService(logger log.Logger,
	uc userV1.UserServiceClient,
	tc tokenV1.TokenServiceClient) AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "admin-srv/service/authentication"))
	return &authenticationService{
		log: l,
		uc:  uc,
		tc:  tc,
	}
}

// Login returns a login reply
func (s *authenticationService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	user, err := s.uc.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	tokenReply, err := s.tc.GenerateToken(ctx, &tokenV1.GenerateTokenRequest{
		TokenInfo: &authn.Claims{
			Subject:  user.GetUserName(),
			UserId:   user.GetId(),
			UserName: user.GetUserName(),
			Platform: tokenV1.TokenPlatform_TOKEN_PLATFORM_WEB.String(),
			NickName: user.GetNickName(),
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		Id:       user.GetId(),
		UserName: user.GetUserName(),
		Token:    tokenReply.GetToken(),
	}, nil
}

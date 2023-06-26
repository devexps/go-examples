package service

import (
	"context"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/backend_api/v1"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/authn"
	tokenV1 "github.com/devexps/go-examples/micro-blog/api/gen/go/token_service/v1"
	userV1 "github.com/devexps/go-examples/micro-blog/api/gen/go/user_service/v1"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/authn/engine"
	"github.com/golang/protobuf/ptypes/empty"
)

// AuthenticationService interface
type AuthenticationService interface {
	v1.AuthenticationServiceServer
}

// authenticationService is an authentication service.
type authenticationService struct {
	v1.UnimplementedAuthenticationServiceServer

	log *log.Helper
	uc  userV1.UserServiceClient
	tc  tokenV1.TokenServiceClient
}

// NewAuthenticationService new an authentication service.
func NewAuthenticationService(logger log.Logger,
	uc userV1.UserServiceClient,
	tc tokenV1.TokenServiceClient,
) AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "backend_api/service/authentication"))
	return &authenticationService{
		log: l,
		uc:  uc,
		tc:  tc,
	}
}

// Login .
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

// Logout .
func (s *authenticationService) Logout(ctx context.Context, req *v1.LogoutRequest) (*empty.Empty, error) {
	claims, _ := engine.FromContext(ctx)
	_, err := s.tc.RemoveToken(ctx, &tokenV1.RemoveTokenRequest{
		TokenInfo: claims.(*authn.Claims),
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *authenticationService) GetMe(ctx context.Context, req *empty.Empty) (*userV1.User, error) {
	claims, _ := engine.FromContext(ctx)
	return s.uc.GetUser(ctx, &userV1.GetUserRequest{
		Id: claims.(*authn.Claims).GetUserId(),
	})
}

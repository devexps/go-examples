package data

import (
	"context"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	tokenV1 "github.com/devexps/go-examples/micro-blog/api/gen/go/token_service/v1"
	userV1 "github.com/devexps/go-examples/micro-blog/api/gen/go/user_service/v1"
	"github.com/devexps/go-examples/micro-blog/pkg/bootstrap"
	"github.com/devexps/go-examples/micro-blog/pkg/middleware"
	"github.com/devexps/go-examples/micro-blog/pkg/service"
	"github.com/devexps/go-micro/v2/log"
	authnEngine "github.com/devexps/go-micro/v2/middleware/authn/engine"
	"github.com/devexps/go-micro/v2/registry"
)

// Data .
type Data struct {
	log *log.Helper

	authenticator authnEngine.Authenticator

	userClient  userV1.UserServiceClient
	tokenClient tokenV1.TokenServiceClient
}

// NewData .
func NewData(cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator,
	userClient userV1.UserServiceClient,
	tokenClient tokenV1.TokenServiceClient,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{
		log:           l,
		authenticator: authenticator,
		userClient:    userClient,
		tokenClient:   tokenClient,
	}, cleanup, nil
}

// NewDiscovery .
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
}

// NewAuthenticator .
func NewAuthenticator(tokenClient tokenV1.TokenServiceClient) authnEngine.Authenticator {
	return middleware.NewAuthenticator(tokenClient)
}

// NewUserServiceClient .
func NewUserServiceClient(cfg *conf.Bootstrap, r registry.Discovery) userV1.UserServiceClient {
	return userV1.NewUserServiceClient(bootstrap.CreateGrpcClient(cfg, context.Background(), r, service.UserService))
}

// NewTokenServiceClient .
func NewTokenServiceClient(cfg *conf.Bootstrap, r registry.Discovery) tokenV1.TokenServiceClient {
	return tokenV1.NewTokenServiceClient(bootstrap.CreateGrpcClient(cfg, context.Background(), r, service.TokenService))
}

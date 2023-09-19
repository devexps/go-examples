package data

import (
	"context"
	"github.com/devexps/go-examples/k8s/api/gen/go/common/conf"
	tokenV1 "github.com/devexps/go-examples/k8s/api/gen/go/token-srv/v1"
	userV1 "github.com/devexps/go-examples/k8s/api/gen/go/user-srv/v1"
	"github.com/devexps/go-examples/k8s/pkg/bootstrap"
	"github.com/devexps/go-examples/k8s/pkg/service"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"
)

// Data .
type Data struct {
	log *log.Helper

	userClient  userV1.UserServiceClient
	tokenClient tokenV1.TokenServiceClient
}

// NewData .
func NewData(cfg *conf.Bootstrap, logger log.Logger,
	userClient userV1.UserServiceClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "admin-srv/data"))
	d := &Data{
		log:        l,
		userClient: userClient,
	}
	return d, func() {
		l.Info("closing the data resources")
	}, nil
}

// NewDiscovery .
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewK8sRegistry(cfg.Registry)
}

// NewUserServiceClient .
func NewUserServiceClient(cfg *conf.Bootstrap, r registry.Discovery) userV1.UserServiceClient {
	return userV1.NewUserServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.UserService, cfg.Client.Grpc.GetTimeout()))
}

// NewTokenServiceClient .
func NewTokenServiceClient(cfg *conf.Bootstrap, r registry.Discovery) tokenV1.TokenServiceClient {
	return tokenV1.NewTokenServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.TokenService, cfg.Client.Grpc.GetTimeout()))
}

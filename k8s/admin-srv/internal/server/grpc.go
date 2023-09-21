package server

import (
	"github.com/devexps/go-examples/k8s/admin-srv/internal/service"
	v1 "github.com/devexps/go-examples/k8s/api/gen/go/admin-srv/v1"
	"github.com/devexps/go-examples/k8s/api/gen/go/common/conf"
	"github.com/devexps/go-examples/k8s/pkg/bootstrap"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	authenticationSvc service.AuthenticationService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))
	v1.RegisterAuthenticationServiceServer(srv, authenticationSvc)
	return srv
}

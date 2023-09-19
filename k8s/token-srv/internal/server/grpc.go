package server

import (
	"github.com/devexps/go-examples/k8s/api/gen/go/common/conf"
	v1 "github.com/devexps/go-examples/k8s/api/gen/go/token-srv/v1"
	"github.com/devexps/go-examples/k8s/pkg/bootstrap"
	"github.com/devexps/go-examples/k8s/token-srv/internal/service"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	tokenSvc service.TokenService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))
	v1.RegisterTokenServiceServer(srv, tokenSvc)
	return srv
}

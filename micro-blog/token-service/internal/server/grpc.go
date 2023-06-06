package server

import (
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/grpc"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/token-service/v1"
	"github.com/devexps/go-examples/micro-blog/token-service/internal/service"
	"github.com/devexps/go-examples/micro-blog/pkg/bootstrap"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	greeterSvc service.GreeterService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))
	v1.RegisterGreeterServer(srv, greeterSvc)
	return srv
}

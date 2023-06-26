package server

import (
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/backend_api/v1"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	"github.com/devexps/go-examples/micro-blog/backend_api/internal/service"
	"github.com/devexps/go-examples/micro-blog/pkg/bootstrap"
	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap,
	middlewares []middleware.Middleware,
	authenticationSvc service.AuthenticationService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, middlewares...)
	v1.RegisterAuthenticationServiceServer(srv, authenticationSvc)
	return srv
}

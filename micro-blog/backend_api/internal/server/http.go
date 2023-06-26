package server

import (
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/backend_api/v1"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	"github.com/devexps/go-examples/micro-blog/backend_api/internal/service"
	"github.com/devexps/go-examples/micro-blog/pkg/bootstrap"
	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *conf.Bootstrap,
	middlewares []middleware.Middleware,
	authenticationSvc service.AuthenticationService,
) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, middlewares...)
	v1.RegisterAuthenticationServiceHTTPServer(srv, authenticationSvc)
	return srv
}

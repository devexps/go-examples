package server

import (
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/http"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/token-service/v1"
	"github.com/devexps/go-examples/micro-blog/token-service/internal/service"
	"github.com/devexps/go-examples/micro-blog/pkg/bootstrap"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *conf.Bootstrap, logger log.Logger,
	greeterSvc service.GreeterService,
) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, logging.Server(logger))
	v1.RegisterGreeterHTTPServer(srv, greeterSvc)
	return srv
}

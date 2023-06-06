package server

import (
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/http"
	"helloworld/api/gen/go/common/conf"
	v1 "helloworld/api/gen/go/helloworld/v1"
	"helloworld/helloworld/internal/service"
	"helloworld/pkg/bootstrap"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *conf.Bootstrap, logger log.Logger,
	greeterSvc service.GreeterService,
) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, logging.Server(logger))
	v1.RegisterGreeterHTTPServer(srv, greeterSvc)
	return srv
}

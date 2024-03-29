package bootstrap

import (
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	"github.com/gorilla/handlers"

	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/metrics"
	"github.com/devexps/go-micro/v2/middleware/ratelimiter"
	"github.com/devexps/go-micro/v2/middleware/recovery"
	"github.com/devexps/go-micro/v2/middleware/tracing"
	"github.com/devexps/go-micro/v2/middleware/validate"
	microHttp "github.com/devexps/go-micro/v2/transport/http"
)

// CreateHttpServer creates an HTTP server
func CreateHttpServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *microHttp.Server {
	var opts = []microHttp.ServerOption{
		microHttp.Filter(handlers.CORS(
			handlers.AllowedHeaders(cfg.Server.Http.Cors.Headers),
			handlers.AllowedMethods(cfg.Server.Http.Cors.Methods),
			handlers.AllowedOrigins(cfg.Server.Http.Cors.Origins),
		)),
	}

	var ms []middleware.Middleware
	if cfg.Server.Http.Middleware != nil {
		if cfg.Server.Http.Middleware.GetEnableRecovery() {
			ms = append(ms, recovery.Recovery())
		}
		if cfg.Server.Http.Middleware.GetEnableTracing() {
			ms = append(ms, tracing.Server())
		}
		if cfg.Server.Http.Middleware.GetEnableValidate() {
			ms = append(ms, validate.Validator())
		}
		if cfg.Server.Http.Middleware.GetEnableRateLimiter() {
			ms = append(ms, ratelimiter.Server())
		}
		if cfg.Server.Http.Middleware.GetEnableMetrics() {
			ms = append(ms, metrics.Server(withMetricRequests(), withMetricHistogram()))
		}
	}
	ms = append(ms, m...)
	opts = append(opts, microHttp.Middleware(ms...))

	if cfg.Server.Http.Network != "" {
		opts = append(opts, microHttp.Network(cfg.Server.Http.Network))
	}
	if cfg.Server.Http.Addr != "" {
		opts = append(opts, microHttp.Address(cfg.Server.Http.Addr))
	}
	if cfg.Server.Http.Timeout != nil {
		opts = append(opts, microHttp.Timeout(cfg.Server.Http.Timeout.AsDuration()))
	}
	srv := microHttp.NewServer(opts...)
	if cfg.Server.Http.Middleware.GetEnableMetrics() {
		handleMetrics(srv)
	}
	return srv
}

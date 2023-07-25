package bootstrap

import (
	"context"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/metrics"
	"github.com/devexps/go-micro/v2/middleware/recovery"
	"github.com/devexps/go-micro/v2/middleware/tracing"
	"github.com/devexps/go-micro/v2/registry"
	microGrpc "github.com/devexps/go-micro/v2/transport/grpc"
	"google.golang.org/grpc"
	"time"
)

const defaultTimeout = 5 * time.Second

// CreateGrpcClient creates a GRPC client
func CreateGrpcClient(cfg *conf.Bootstrap, ctx context.Context, r registry.Discovery, serviceName string, m ...middleware.Middleware) grpc.ClientConnInterface {
	timeout := defaultTimeout
	if cfg.Client.Grpc.Timeout != nil {
		timeout = cfg.Client.Grpc.Timeout.AsDuration()
	}
	endpoint := "discovery:///" + serviceName

	var ms []middleware.Middleware
	if cfg.Client.Grpc.Middleware != nil {
		if cfg.Client.Grpc.Middleware.GetEnableRecovery() {
			ms = append(ms, recovery.Recovery())
		}
		if cfg.Client.Grpc.Middleware.GetEnableTracing() {
			ms = append(ms, tracing.Client())
		}
	}
	ms = append(ms, m...)

	conn, err := microGrpc.DialInsecure(
		ctx,
		microGrpc.WithEndpoint(endpoint),
		microGrpc.WithDiscovery(r),
		microGrpc.WithTimeout(timeout),
		microGrpc.WithMiddleware(ms...),
	)
	if err != nil {
		log.Fatalf("dial grpc client [%s] failed: %s", serviceName, err.Error())
	}
	return conn
}

// CreateGrpcServer Create GRPC server
func CreateGrpcServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *microGrpc.Server {
	var opts []microGrpc.ServerOption

	var ms []middleware.Middleware
	if cfg.Server.Grpc.Middleware != nil {
		if cfg.Server.Grpc.Middleware.GetEnableRecovery() {
			ms = append(ms, recovery.Recovery())
		}
		if cfg.Server.Grpc.Middleware.GetEnableTracing() {
			ms = append(ms, tracing.Server())
		}
		if cfg.Server.Http.Middleware.GetEnableMetrics() {
			ms = append(ms, metrics.Server(withMetricRequests(), withMetricHistogram()))
		}
	}
	ms = append(ms, m...)
	opts = append(opts, microGrpc.Middleware(ms...))

	if cfg.Server.Grpc.Network != "" {
		opts = append(opts, microGrpc.Network(cfg.Server.Grpc.Network))
	}
	if cfg.Server.Grpc.Addr != "" {
		opts = append(opts, microGrpc.Address(cfg.Server.Grpc.Addr))
	}
	if cfg.Server.Grpc.Timeout != nil {
		opts = append(opts, microGrpc.Timeout(cfg.Server.Grpc.Timeout.AsDuration()))
	}
	return microGrpc.NewServer(opts...)
}

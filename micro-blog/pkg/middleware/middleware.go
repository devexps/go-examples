package middleware

import (
	"context"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware"
	midAuthn "github.com/devexps/go-micro/v2/middleware/authn"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/middleware/selector"
	"github.com/devexps/go-micro/v2/transport"
)

// NewMiddlewares .
func NewMiddlewares(opts ...Option) []middleware.Middleware {
	ops := &options{}
	for _, o := range opts {
		o(ops)
	}
	var ms []middleware.Middleware
	if ops.logger != nil {
		ms = append(ms, logging.Server(ops.logger))
	}
	var sms []middleware.Middleware
	if ops.authenticator != nil {
		sms = append(sms, midAuthn.Server(ops.authenticator))
	}
	sms = append(sms, Server())
	sec := selector.Server(sms...)
	if ops.matcher != nil {
		sec = sec.Match(ops.matcher)
	}
	ms = append(ms, sec.Build())
	return ms
}

// Server .
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}
			log.Info(tr)
			authnClaims, ok := midAuthn.FromContext(ctx)
			if !ok {
				return handler(ctx, req)
			}
			log.Info(authnClaims)
			return handler(ctx, req)
		}
	}
}

package metadata

import (
	"context"
	"github.com/devexps/go-micro/v2/log"
	microMetadata "github.com/devexps/go-micro/v2/metadata"
	"github.com/devexps/go-micro/v2/middleware"
	microMidMetadata "github.com/devexps/go-micro/v2/middleware/metadata"
)

// Server is middleware server-side metadata.
func Server() middleware.Middleware {
	return microMidMetadata.Server(microMidMetadata.WithPropagatedPrefix("x-"))
}

// Client is middleware client-side metadata.
func Client() middleware.Middleware {
	return microMidMetadata.Client(microMidMetadata.WithPropagatedPrefix("x-"))
}

// RemoteAddr returns a ip valuer.
func RemoteAddr() log.Valuer {
	return func(ctx context.Context) interface{} {
		if md, ok := microMetadata.FromServerContext(ctx); ok {
			if ip := md.Get("X-Forwarded-For"); ip != "" {
				return ip
			}
			if ip := md.Get("X-Real-IP"); ip != "" {
				return ip
			}
		}
		return ""
	}
}

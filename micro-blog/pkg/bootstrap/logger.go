package bootstrap

import (
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	"github.com/devexps/go-examples/micro-blog/pkg/middleware/metadata"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/tracing"
	"os"
)

type LoggerType string

const (
	LoggerTypeStd    LoggerType = "std"
	LoggerTypeLogrus LoggerType = "logrus"
	LoggerTypeZap    LoggerType = "zap"
)

// NewLoggerProvider creates a new logger provider
func NewLoggerProvider(cfg *conf.Logger, serviceInfo *ServiceInfo) log.Logger {
	l := NewLogger(cfg)

	return log.With(
		l,
		"service.id", serviceInfo.Id,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
		"client_ip", metadata.RemoteAddr(),
	)
}

// NewLogger creates a new logger
func NewLogger(cfg *conf.Logger) log.Logger {
	if cfg == nil {
		return NewStdLogger()
	}

	switch LoggerType(cfg.Type) {
	default:
		fallthrough
	case LoggerTypeStd:
		return NewStdLogger()
	}
}

// NewStdLogger creates a new logger - Micro built-in, console output
func NewStdLogger() log.Logger {
	return log.NewStdLogger(os.Stdout)
}

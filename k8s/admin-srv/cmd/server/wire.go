//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/devexps/go-examples/k8s/admin-srv/internal/data"
	"github.com/devexps/go-examples/k8s/admin-srv/internal/server"
	"github.com/devexps/go-examples/k8s/admin-srv/internal/service"
	"github.com/devexps/go-examples/k8s/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"
	"github.com/google/wire"
)

// initApp init micro application.
func initApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*micro.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}

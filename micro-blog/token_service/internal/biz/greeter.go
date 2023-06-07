package biz

import (
	"context"
	"github.com/devexps/go-examples/micro-blog/token_service/internal/data"
	"github.com/devexps/go-micro/v2/log"
)

// GreeterUseCase is a Greater use case interface.
type GreeterUseCase interface {
	CreateGreeter(context.Context, *data.Greeter) (*data.Greeter, error)
}

// greeterUseCase is a Greeter use case.
type greeterUseCase struct {
	repo data.GreeterRepo
	log  *log.Helper
}

// NewGreeterUseCase new a Greeter use case.
func NewGreeterUseCase(repo data.GreeterRepo, logger log.Logger) GreeterUseCase {
	return &greeterUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *greeterUseCase) CreateGreeter(ctx context.Context, g *data.Greeter) (*data.Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

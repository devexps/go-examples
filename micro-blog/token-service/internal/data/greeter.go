package data

import (
	"context"
	"github.com/devexps/go-micro/v2/log"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo interface.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
}

// greeterRepo is a Greeter repository
type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save stores a greeter to db
func (r *greeterRepo) Save(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

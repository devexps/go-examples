package service

import (
	"context"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/token_service/v1"
	"github.com/devexps/go-examples/micro-blog/token_service/internal/biz"
	"github.com/devexps/go-examples/micro-blog/token_service/internal/data"
)

type GreeterService interface {
	v1.GreeterServer
}

// GreeterService is a greeter service.
type greeterService struct {
	v1.UnimplementedGreeterServer

	uc biz.GreeterUseCase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc biz.GreeterUseCase) GreeterService {
	return &greeterService{uc: uc}
}

// SayHello implements token_service.GreeterServer.
func (s *greeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &data.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

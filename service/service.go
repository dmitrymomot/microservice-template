package service

import (
	"context"

	"github.com/dmitrymomot/microservice-template/pb/example"
)

type service struct {
	example.UnimplementedServiceServer
}

// NewService is a factory function, returns a new instance of the service structure
func NewService() example.ServiceServer {
	return &service{}
}

func (*service) Call(ctx context.Context, req *example.Req) (*example.Resp, error) {
	return &example.Resp{Str: "Received: " + req.GetStr()}, nil
}

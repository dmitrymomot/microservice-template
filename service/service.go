package service

import (
	"context"

	"github.com/dmitrymomot/microservice-template/pb/examplesrv"
)

type service struct {
	examplesrv.UnimplementedServiceServer
}

// NewService is a factory function, returns a new instance of the service structure
func NewService() examplesrv.ServiceServer {
	return &service{}
}

func (*service) Call(ctx context.Context, req *examplesrv.Req) (*examplesrv.Resp, error) {
	return &examplesrv.Resp{Str: "Received: " + req.GetStr()}, nil
}

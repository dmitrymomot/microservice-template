package service

import (
	"context"

	"github.com/dmitrymomot/microservice-template/pb"
)

type service struct {
	pb.UnimplementedServiceServer
}

// NewService is a factory function, returns a new instance of the service structure
func NewService() pb.ServiceServer {
	return &service{}
}

func (*service) Call(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{Str: "Received: " + req.GetStr()}, nil
}

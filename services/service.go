package services

import "github.com/dmitrymomot/microservice-template/proto/microservice_template"

type service struct {
	microservice_template.UnimplementedServiceServer
}

// NewService is a factory function, returns a new instance of the service structure
func NewService() microservice_template.ServiceServer {
	return &service{}
}

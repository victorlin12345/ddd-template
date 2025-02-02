package presentation

import (
	"context"
	"fmt"

	app "github.com/victorlin12345/ddd-template/internal/application"
	grpcserver "github.com/victorlin12345/ddd-template/internal/infrastructure/grpc_server"
	hellopb "github.com/victorlin12345/ddd-template/internal/pkg/gen"
)

type HelloController struct {
	hellopb.UnimplementedMyServiceServer
	orderProcess app.OrderProcess
}

func NewHelloController(server *grpcserver.GrpcServer, orderProcess app.OrderProcess) *HelloController {
	ctr := &HelloController{}
	hellopb.RegisterMyServiceServer(server, ctr)
	ctr.orderProcess = orderProcess

	return ctr
}

func (c *HelloController) SayHello(context.Context, *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Println("Hello")
	if err := c.orderProcess.CreateOrder(); err != nil {
		return nil, err
	}

	return &hellopb.HelloResponse{
		Message: "Hello",
	}, nil
}

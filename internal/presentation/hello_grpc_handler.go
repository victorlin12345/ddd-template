package presentation

import (
	"context"
	"fmt"

	"github.com/victorlin12345/ddd-template/internal/application"
	grpcserver "github.com/victorlin12345/ddd-template/internal/infrastructure/grpc_server"
	hellopb "github.com/victorlin12345/ddd-template/internal/pkg/gen"
)

type HelloController struct {
	hellopb.UnimplementedMyServiceServer
	orderManger *application.OrderManger
}

func NewHelloController(server *grpcserver.GrpcServer, orderManger *application.OrderManger) *HelloController {
	ctr := &HelloController{
		orderManger: orderManger,
	}
	hellopb.RegisterMyServiceServer(server, ctr)

	return ctr
}

func (c *HelloController) SayHello(context.Context, *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Println("Hello")

	if err := c.orderManger.CreateOrder(); err != nil {
		Wrap(ErrDatabase, "failed to save order", err)
		return nil, err
	}

	return &hellopb.HelloResponse{
		Message: "Hello",
	}, nil
}

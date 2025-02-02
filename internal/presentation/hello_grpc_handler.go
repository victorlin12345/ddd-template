package presentation

import (
	"context"
	"fmt"

	grpcserver "github.com/victorlin12345/ddd-template/internal/infrastructure/grpc_server"
	hellopb "github.com/victorlin12345/ddd-template/internal/pkg/gen"
)

type HelloController struct {
	hellopb.UnimplementedMyServiceServer
}

func NewHelloController(server *grpcserver.GrpcServer) *HelloController {
	ctr := &HelloController{}
	hellopb.RegisterMyServiceServer(server, ctr)

	return ctr
}

func (c *HelloController) SayHello(context.Context, *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Println("Hello")
	return &hellopb.HelloResponse{
		Message: "Hello",
	}, nil
}

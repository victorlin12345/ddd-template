package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	*grpc.Server
}

func NewGrpcServer(lc fx.Lifecycle) *GrpcServer {

	srv := &GrpcServer{
		Server: grpc.NewServer(),
	}

	return srv
}

func (srv *GrpcServer) Start(ctx context.Context) error {

	// 監聽指定的端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 啟動服務
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v",
				err)
		}
	}()

	return nil
}

func (srv *GrpcServer) Stop(ctx context.Context) error {
	fmt.Println("grpc server stop")
	srv.GracefulStop()
	return nil
}

package infrastructure

import (
	grpcserver "github.com/victorlin12345/ddd-template/internal/infrastructure/grpc_server"
	"go.uber.org/fx"
)

var Module = fx.Module("infrastructure",
	fx.Provide(grpcserver.NewGrpcServer),
)

package infrastructure

import (
	grpcserver "github.com/victorlin12345/ddd-template/internal/infrastructure/grpc_server"
	"github.com/victorlin12345/ddd-template/internal/infrastructure/mongo"
	"go.uber.org/fx"
)

var Module = fx.Module("infrastructure",
	fx.Provide(mongo.NewMongoClient),
	fx.Provide(grpcserver.NewGrpcServer),
)

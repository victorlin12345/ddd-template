package presentation

import "go.uber.org/fx"

var Module = fx.Module("presentation",
	fx.Provide(NewGrpcRouter),
	fx.Provide(NewHelloController),
)

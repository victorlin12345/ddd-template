package repository

import "go.uber.org/fx"

var Module = fx.Module("repostioty",
	fx.Provide(NewOrderMongoRepo),
	fx.Provide(NewPaymentRepo),
)

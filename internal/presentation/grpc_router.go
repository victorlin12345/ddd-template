package presentation

type GrpcRouter struct {
	helloCtr *HelloController
}

func NewGrpcRouter(helloCtr *HelloController) *GrpcRouter {
	return &GrpcRouter{helloCtr: helloCtr}
}

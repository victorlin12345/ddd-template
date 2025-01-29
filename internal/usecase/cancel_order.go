package application

import (
	"context"

	"github.com/victorlin12345/ddd-template/internal/dto"
)

func CancelOrder(ctx context.Context, params dto.CancelOrderParams) (*dto.CancelOrderResponse, error) {

	// req
	// auth

	// getTx
	// get order

	// get grpc client

	// save order

	// use order do domain logic

	// payment

	// save order

	// service.OrderService.CancelOrder()
	// service.OrderPaymentService.DoSomething()
	// resp
	return &dto.CancelOrderResponse{}, nil
}

package application

import (
	"context"
	"fmt"

	"github.com/victorlin12345/ddd-template/internal/domain/order"
)

type OrderManger struct {
	orderRepo order.Repository
}

func NewOrderManger(orderRepo order.Repository) *OrderManger {
	return &OrderManger{
		orderRepo: orderRepo,
	}
}

func (p *OrderManger) CreateOrder() error {
	ctx := context.Background()
	order := order.NewOrder()
	order, err := p.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return err
	}
	fmt.Printf("[OrderManger] Create order: %+v", order)
	return nil
}

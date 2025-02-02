package application

import (
	"context"
	"fmt"

	"github.com/victorlin12345/ddd-template/internal/domain/order"
	"github.com/victorlin12345/ddd-template/internal/domain/payment"
	"github.com/victorlin12345/ddd-template/internal/domain/service"
	"github.com/victorlin12345/ddd-template/internal/repository"
)

type OrderProcess struct {
	orderRepo       order.Repository
	paymentRepo     payment.Repository
	orderPaymentSvc service.OrderPaymentService
}

func NewOrderProcess() *OrderProcess {
	return &OrderProcess{
		orderRepo:       repository.NewOrderRepo(),
		paymentRepo:     repository.NewPaymentRepo(),
		orderPaymentSvc: service.OrderPaymentService{},
	}
}

func (p *OrderProcess) CreateOrder() error {
	ctx := context.Background()
	return p.orderRepo.Save(ctx, order.NewOrder())
}

func (p *OrderProcess) CancelOrder(orderID int64) error {
	ctx := context.Background()
	o, err := p.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to get order by id: %w", err)
	}
	o.CancelOrder()
	return p.orderRepo.Save(ctx, o)
}

func (p *OrderProcess) PaymentExpired(orderID int64) error {
	ctx := context.Background()
	o, err := p.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to get order by id: %w", err)
	}
	payment, err := p.paymentRepo.GetByOrderID(orderID)
	if err != nil {
		return fmt.Errorf("failed to get payment by order id: %w", err)
	}
	if err := p.orderPaymentSvc.PaymentExpired(o, payment); err != nil {
		return fmt.Errorf("failed to process payment expired: %w", err)
	}
	if err := p.orderRepo.Save(ctx, o); err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}
	if err := p.paymentRepo.Save(payment); err != nil {
		return fmt.Errorf("failed to save payment: %w", err)
	}
	return nil
}

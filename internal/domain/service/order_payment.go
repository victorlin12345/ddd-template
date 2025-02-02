package service

import (
	"github.com/victorlin12345/ddd-template/internal/domain/order"
	"github.com/victorlin12345/ddd-template/internal/domain/payment"
)

// business logic need to be aggregated purpose
type OrderPaymentService struct {
}

func (srv *OrderPaymentService) PaymentExpired(order *order.Order, payment *payment.Payment) error {
	order.CancelOrder()
	payment.CancelPayment()
	return nil
}

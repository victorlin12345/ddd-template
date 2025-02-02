package repository

import "github.com/victorlin12345/ddd-template/internal/domain/payment"

type PaymentMongoRepo struct{}

// GetByOrderID implements payment.Repository.
func (p *PaymentMongoRepo) GetByOrderID(id int64) (*payment.Payment, error) {
	panic("unimplemented")
}

// Save implements payment.Repository.
func (p *PaymentMongoRepo) Save(payment *payment.Payment) error {
	panic("unimplemented")
}

func NewPaymentRepo() payment.Repository {
	return &PaymentMongoRepo{}
}

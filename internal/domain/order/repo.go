package order

import "context"

type Repository interface {
	GetByID(ctx context.Context, id int64) (*Order, error)
	Save(ctx context.Context, order *Order) error
	CreateOrder(ctx context.Context, order *Order) (*Order, error)
}

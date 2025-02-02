package payment

type Repository interface {
	GetByOrderID(id int64) (*Payment, error)
	Save(payment *Payment) error
}

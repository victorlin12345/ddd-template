package payment

type Repository interface {
	GetByID(id int64) (*Payment, error)
	Save(payment *Payment) error
}

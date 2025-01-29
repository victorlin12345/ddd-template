package order

type Repository interface {
	GetByID(id int64) (*Order, error)
	Save(payment *Order) error
}

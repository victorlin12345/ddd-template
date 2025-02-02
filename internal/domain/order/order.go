package order

type Status string

const (
	OrderStatusCreated  Status = "CREATED"
	OrderStatusCanceled Status = "CANCELLED"
)

// Entity
type Order struct {
	ID     int64
	Status Status
}

func NewOrder() *Order {
	return &Order{
		Status: OrderStatusCreated,
	}
}

// Cancel logic
func (o *Order) CancelOrder() {
	o.Status = OrderStatusCanceled
}

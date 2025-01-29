package order

// Entity
type Order struct {
}

func NewOrder() *Order {
	return &Order{}
}

// Cancel logic
func (o *Order) CancelOrder() {
}

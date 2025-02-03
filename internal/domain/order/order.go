package order

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	OrderStatusCreated  Status = "CREATED"
	OrderStatusCanceled Status = "CANCELLED"
)

// Entity
type Order struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Status Status             `bson:"status"`
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

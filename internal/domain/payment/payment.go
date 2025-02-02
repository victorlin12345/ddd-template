package payment

type Status string

const (
	PaymentStatusCreated  Status = "CREATED"
	PaymentStatusCanceled Status = "CANCELLED"
)

type Payment struct {
	OrderID int64
	Status  Status
}

func NewPayment(orderID int64) *Payment {
	return &Payment{
		OrderID: orderID,
		Status:  PaymentStatusCreated,
	}
}

func (p *Payment) CancelPayment() {
	p.Status = PaymentStatusCanceled
}

package dto

type CancelOrderParams struct{}

type CancelOrderResponse struct{}

func (p *CancelOrderParams) Validate() error {
	return nil
}

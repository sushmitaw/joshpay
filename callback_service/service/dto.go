package service

type CallBackRequest struct {
	PaymentID string `json:"payment_id"`
	Status    bool   `json:"status"`
}

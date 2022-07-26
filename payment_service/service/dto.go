package service

type PaymentRequest struct {
	CreditCardToken string  `json:"credit_card_token"`
	Amount          float64 `json:"amount"`
}

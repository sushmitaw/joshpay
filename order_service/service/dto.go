package service

type CreateOrderRequest struct {
	Items      []string `json:"items"`
	Amount     float64  `json:"amount"`
	CustomerID string   `json:"customer_id"`
}

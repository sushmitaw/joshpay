package service

type CustomerDetailsRequest struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	MobileNumber string `json:"mobile_number"`
}

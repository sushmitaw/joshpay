package service

func InitiatePaymentService(req PaymentRequest) error {
	//simulate payment success
	go checkIfSuspicious(req)
	return nil
}

func checkIfSuspicious(req PaymentRequest) {
	//get history from database
	//history := []fraud_detection_service.PaymentRequest{}
	//
	//isSuspicious := fraud_detection_service.CheckTransaction(req.Amount, req.CreditCardToken, history)
	//if isSuspicious {
	//	//add in db for further investigation
	//}

}

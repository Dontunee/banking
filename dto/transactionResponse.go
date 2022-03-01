package dto

type TransactionResponse struct {
	Balance       float64 `json:"balance"`
	TransactionId string  `json:"transaction_id"`
}

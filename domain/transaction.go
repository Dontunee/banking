package domain

import (
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/errs"
)

type Transaction struct {
	TransactionType string
	Amount          float64
	TransactionID   string
	TransactionDate string
	AccountId       string
}

func (transaction Transaction) ToProcessTransactionDto(updatedBalance float64, transactionId string) *dto.TransactionResponse {
	transactionResponse := dto.TransactionResponse{
		Balance:       updatedBalance,
		TransactionId: transactionId,
	}
	return &transactionResponse
}

type ITransactionRepository interface {
	Create(transaction Transaction) (*Transaction, *errs.AppError)
}

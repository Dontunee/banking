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

func (transaction *Transaction) Init(transactionType string, amount float64, date string, accountId string) {
	transaction.TransactionType = transactionType
	transaction.Amount = amount
	transaction.TransactionDate = date
	transaction.AccountId = accountId
}

func (transaction Transaction) ToProcessTransactionDto(updatedBalance float64, transactionId string) *dto.TransactionResponse {
	transactionResponse := dto.TransactionResponse{
		Balance:       updatedBalance,
		TransactionId: transactionId,
	}
	return &transactionResponse
}

func (transaction Transaction) IsTransactionTypeWithdrawal() bool {
	if transaction.TransactionType == "withdrawal" {
		return true
	}
	return false
}

type ITransactionRepository interface {
	Create(transaction Transaction) (*Transaction, *errs.AppError)
}

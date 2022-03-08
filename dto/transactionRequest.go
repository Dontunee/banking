package dto

import (
	"github.com/Dontunee/banking/errs"
)

type TransactionRequest struct {
	Amount          float64 `json:"amount"`
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
}

func (request TransactionRequest) Validate() (*errs.AppError, bool) {
	if !request.IsTransactionTypeWithdrawal() && !request.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be withdrawal or deposit"), false
	}
	if request.Amount < 0 {
		return errs.NewValidationError("Amount can only be a positive number"), false
	}
	return nil, true
}

func (request TransactionRequest) CanWithdraw(amount float64, balance float64) (bool, *errs.AppError) {
	if amount <= 0 {
		return false, errs.NewValidationError("Amount to withdraw is less than zero")
	}
	if amount > balance {
		return false, errs.NewValidationError("Amount to withdraw greater than balance")
	}
	return true, nil
}

func (request TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return request.TransactionType == "withdrawal"
}

func (request TransactionRequest) IsTransactionTypeDeposit() bool {
	return request.TransactionType == "deposit"
}

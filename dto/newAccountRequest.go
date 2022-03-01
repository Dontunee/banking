package dto

import (
	"github.com/Dontunee/banking/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (request NewAccountRequest) IsValid() (*errs.AppError, bool) {
	if request.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000"), false
	}
	if strings.ToLower(request.AccountType) != "saving" || strings.ToLower(request.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving"), false
	}
	return nil, true
}

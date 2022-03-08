package domain

import (
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/errs"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      bool
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/Dontunee/banking/domain IAccountRepository
type IAccountRepository interface {
	Create(account Account) (*Account, *errs.AppError)
	UpdateBalance(amount float64, accountId string) (*errs.AppError, bool)
	GetBalance(accountId string) (balance float64, appError *errs.AppError)
}

func (account Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: account.AccountID}
}

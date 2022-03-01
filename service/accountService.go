package service

import (
	"github.com/Dontunee/banking/domain"
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/errs"
	"time"
)

type IAccountService interface {
	CreateAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type AccountService struct {
	repository domain.IAccountRepository
}

func (accountService AccountService) CreateAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err, valid := request.IsValid()
	if !valid {
		return nil, err
	}

	account := domain.Account{
		AccountID:   "",
		CustomerID:  request.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      true,
	}
	newAccount, error := accountService.repository.Create(account)
	if error != nil {
		return nil, error
	}
	responseDto := newAccount.ToNewAccountResponseDto()
	return &responseDto, nil

}

func NewAccountService(repository domain.IAccountRepository) AccountService {
	return AccountService{repository: repository}
}

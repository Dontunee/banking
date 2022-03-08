package service

import (
	"github.com/Dontunee/banking/domain"
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/errs"
	"time"
)

type ITransactionService interface {
	ProcessTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type TransactionService struct {
	accountRepo     domain.IAccountRepository
	transactionRepo domain.ITransactionRepository
}

func (transactionService TransactionService) ProcessTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err, valid := request.Validate()
	if !valid {
		return nil, errs.NewValidationError("request is not valid:" + err.Message)
	}
	balance, err := transactionService.accountRepo.GetBalance(request.AccountId)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unable to get balance: " + err.Message)
	}
	transaction := new(domain.Transaction)
	if request.IsTransactionTypeWithdrawal() {
		valid, err = request.CanWithdraw(request.Amount, balance)
		if !valid {
			return nil, errs.NewValidationError("insufficient balance in the account")
		}
	}
	transaction.Init(request.TransactionType, request.Amount, time.Now().Format("2006-01-02 15:04:05"), request.AccountId)
	create, err := transactionService.transactionRepo.Create(*transaction)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unable to create transaction: " + err.Message)
	}
	balance, err = transactionService.accountRepo.GetBalance(request.AccountId)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unable to get balance: " + err.Message)
	}
	responseDto := create.ToProcessTransactionDto(balance, create.TransactionID)
	return responseDto, nil
}

func NewTransactionService(accountRepo domain.IAccountRepository, transactionRepo domain.ITransactionRepository) TransactionService {
	return TransactionService{accountRepo: accountRepo, transactionRepo: transactionRepo}
}

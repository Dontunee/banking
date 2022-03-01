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
	err, valid := request.IsValid()
	if !valid {
		return nil, err
	} else {
		balance, err := transactionService.accountRepo.GetBalance(request.AccountId)
		if err != nil {
			return nil, err
		}
		if request.TransactionType == "withdrawal" {
			if err != nil {
				return nil, err
			}
			valid, err := request.IsWithdrawalValid(request.Amount, balance)
			if !valid {
				return nil, err
			}
			transaction := domain.Transaction{
				TransactionType: request.TransactionType,
				Amount:          request.Amount,
				TransactionID:   "",
				TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
				AccountId:       request.AccountId,
			}
			withdraw, err := transactionService.transactionRepo.Create(transaction)
			if err != nil {
				return nil, err
			}
			//update balance
			err, balanceUpdate := transactionService.accountRepo.UpdateBalance(balance-request.Amount, request.AccountId)
			if balanceUpdate == true {
				responseDto := withdraw.ToProcessTransactionDto(balance-request.Amount, withdraw.TransactionID)
				return responseDto, nil
			} else {
				return nil, err
			}

		} else {
			transaction := domain.Transaction{
				TransactionType: request.TransactionType,
				Amount:          request.Amount,
				TransactionID:   "",
				TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
				AccountId:       request.AccountId,
			}
			deposit, err := transactionService.transactionRepo.Create(transaction)
			if err != nil {
				return nil, err
			}
			responseDto := deposit.ToProcessTransactionDto(balance+request.Amount, deposit.TransactionID)
			return responseDto, nil
		}
	}
}

func NewTransactionService(accountRepo domain.IAccountRepository, transactionRepo domain.ITransactionRepository) TransactionService {
	return TransactionService{accountRepo: accountRepo, transactionRepo: transactionRepo}
}

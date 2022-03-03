package domain

import (
	"github.com/Dontunee/banking/errs"
	"github.com/Dontunee/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (repository TransactionRepositoryDb) Create(transaction Transaction) (*Transaction, *errs.AppError) {
	session, err := repository.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	createSql := "INSERT INTO transactions (account_id,amount,transaction_type,transaction_date) values (?,?,?,?)"
	result, err := repository.client.Exec(createSql, transaction.AccountId, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		logger.Error("Error occurred while creating transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	//update Account balance
	if transaction.IsTransactionTypeWithdrawal() {
		updateSql := "UPDATE accounts SET amount = amount - ? WHERE account_id = ?"
		_, err = repository.client.Exec(updateSql, transaction.Amount, transaction.AccountId)
	} else {
		updateSql := "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
		_, err = repository.client.Exec(updateSql, transaction.Amount, transaction.AccountId)
	}
	if err != nil {
		session.Rollback()
		logger.Error("Error occurred while saving transaction:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = session.Commit()
	if err != nil {
		session.Rollback()
		logger.Error("Error occurred while commiting transaction for bank account:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error occurred while retrieving last inserted id" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	transaction.TransactionID = strconv.FormatInt(id, 10)
	return &transaction, nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}

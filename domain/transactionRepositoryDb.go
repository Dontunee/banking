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
	createSql := "INSERT INTO transactions (account_id,amount,transaction_type,transaction_date) values (?,?,?,?)"

	result, err := repository.client.Exec(createSql, transaction.AccountId, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		logger.Error("Error occurred while creating transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
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

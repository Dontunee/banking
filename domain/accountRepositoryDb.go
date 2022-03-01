package domain

import (
	"database/sql"
	"github.com/Dontunee/banking/errs"
	"github.com/Dontunee/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (repository AccountRepositoryDb) Create(account Account) (*Account, *errs.AppError) {
	createSql := "INSERT INTO accounts (customer_id,opening_date,account_type,amount,status) values (?,?,?,?,?)"
	result, err := repository.client.Exec(createSql, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error occurred while creating account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error occurred while retrieving last inserted id" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	account.AccountID = strconv.FormatInt(id, 10)
	return &account, nil

}

func (repository AccountRepositoryDb) UpdateBalance(amount float64, accountId string) (*errs.AppError, bool) {
	updateSql := "UPDATE accounts SET amount = ? WHERE account_id = ?"
	_, err := repository.client.Exec(updateSql, amount, accountId)
	if err != nil {
		logger.Error("Error occurred while updating account" + err.Error())
		return errs.NewUnexpectedError("Unexpected error from database"), false
	} else {
		return nil, true
	}
}

func (repository AccountRepositoryDb) GetBalance(accountId string) (balance float64, appError *errs.AppError) {
	balanceQuery := "select amount from accounts where account_id = ?"

	err := repository.client.Get(&balance, balanceQuery, &accountId)
	if err != nil {
		logger.Error("Error occurred while querying for balance" + err.Error())
		if err == sql.ErrNoRows {
			return 0, errs.NewNotFoundError("account not found")
		} else {
			return 0, errs.NewUnexpectedError("Unexpected database error ")
		}
	}

	return balance, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

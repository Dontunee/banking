package domain

import (
	"database/sql"
	"github.com/Dontunee/banking/errs"
	"github.com/Dontunee/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (repository CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth,status from customers"
	customers := make([]Customer, 0)

	err := repository.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Error("Error occurred while querying customers" + err.Error())
		return nil, errs.NewNotFoundError("Customers not found")
	}
	return customers, nil
}

func (repository CustomerRepositoryDb) FindCustomerById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth,status from customers where customer_id = ?"
	var customer Customer

	err := repository.client.Get(&customer, customerSql, id)
	if err != nil {
		logger.Error("Error occurred while querying for customer" + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error ")
		}
	}
	return &customer, nil
}

func (repository CustomerRepositoryDb) FindCustomersByStatus(status bool) ([]Customer, *errs.AppError) {
	customersSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	customers := make([]Customer, 0)

	err := repository.client.Select(&customers, customersSql, status)
	if err != nil {
		logger.Error("Error occurred while querying customer table" + err.Error())
	}
	return customers, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}

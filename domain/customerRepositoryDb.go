package domain

import (
	"database/sql"
	"github.com/Dontunee/banking/errs"
	"github.com/Dontunee/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (repository CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth,status from customers"

	rows, err := repository.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error occurred while querying customer table" + err.Error())
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			logger.Error("Error occurred while scanning customers" + err.Error())
			return nil, errs.NewNotFoundError("Customers not found")
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (repository CustomerRepositoryDb) FindCustomerById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth,status from customers where customer_id = ?"

	row := repository.client.QueryRow(customerSql, id)
	var customer Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateOfBirth, &customer.Status)
	if err != nil {
		logger.Error("Error occurred while scanning customer" + err.Error())
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

	rows, err := repository.client.Query(customersSql, status)

	if err != nil {
		logger.Error("Error occurred while querying customer table" + err.Error())
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			logger.Error("Error occurred while scanning customers" + err.Error())
			return nil, errs.NewNotFoundError("Customers not found")
		}

		customers = append(customers, customer)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:password10$@tcp(localhost:3306)/banking")

	if err != nil {
		panic(err)
	}

	//see "Important settings"  section
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}

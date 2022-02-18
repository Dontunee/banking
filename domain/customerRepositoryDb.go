package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (repository CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth,status from customers"

	rows, err := repository.client.Query(findAllSql)
	if err != nil {
		log.Println("Error occurred while querying customer table" + err.Error())
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			log.Println("Error occurred while scanning customers" + err.Error())
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (repository CustomerRepositoryDb) FindCustomerById(id string) (*Customer, error) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth,status from customers where customer_id = ?"

	row := repository.client.QueryRow(customerSql, id)
	var customer Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateOfBirth, &customer.Status)
	if err != nil {
		log.Println("Error occurred while scanning customer" + err.Error())
		return nil, err
	}
	return &customer, nil

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

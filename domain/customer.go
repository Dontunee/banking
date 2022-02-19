package domain

import "github.com/Dontunee/banking/errs"

type Customer struct {
	ID          int
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      bool
}

//secondary port
//repository interface

type ICustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindCustomerById(id string) (*Customer, *errs.AppError)
	FindCustomersByStatus(status bool) ([]Customer, *errs.AppError)
}

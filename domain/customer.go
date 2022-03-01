package domain

import (
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/errs"
)

type Customer struct {
	ID          int `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      bool
}

//secondary port
//repository interface

type ICustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindCustomerById(id string) (*Customer, *errs.AppError)
	FindCustomersByStatus(status bool) ([]Customer, *errs.AppError)
}

func (customer Customer) ParseStatusAsText() string {
	statusAsText := "active"
	if customer.Status == false {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (customer Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		City:        customer.City,
		ZipCode:     customer.ZipCode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.ParseStatusAsText(),
	}
}

func ToSliceDtos(customers []Customer) []dto.CustomerResponse {

	customerDtos := make([]dto.CustomerResponse, 0)
	for _, element := range customers {
		dto := element.ToDto()
		customerDtos = append(customerDtos, dto)
	}

	return customerDtos
}

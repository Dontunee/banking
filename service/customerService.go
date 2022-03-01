package service

import (
	"github.com/Dontunee/banking/domain"
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/errs"
)

type ICustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError)
	GetCustomersByStatus(status bool) ([]dto.CustomerResponse, *errs.AppError)
}

type CustomerService struct {
	repo domain.ICustomerRepository // inject customer repo interface
}

//method in Customer Service class/struct

func (customerService CustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError) {
	allCustomers, err := customerService.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return domain.ToSliceDtos(allCustomers), err
}

func (customerService CustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := customerService.repo.FindCustomerById(id)
	if err != nil {
		return nil, err
	}
	response := customer.ToDto()
	return &response, nil
}

func (customerService CustomerService) GetCustomersByStatus(status bool) ([]dto.CustomerResponse, *errs.AppError) {
	customersByStatus, err := customerService.repo.FindCustomersByStatus(status)
	if err != nil {
		return nil, err
	}

	return domain.ToSliceDtos(customersByStatus), err

}

//constructor method to initiate CustomerService class/struct

func NewCustomerService(repository domain.ICustomerRepository) CustomerService {
	return CustomerService{repository}
}

package service

import (
	"github.com/Dontunee/banking/domain"
	"github.com/Dontunee/banking/errs"
)

type ICustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomerById(id string) (*domain.Customer, *errs.AppError)
	GetCustomersByStatus(status bool) ([]domain.Customer, *errs.AppError)
}

type CustomerService struct {
	repo domain.ICustomerRepository // inject customer repo interface
}

//method in Customer Service class/struct

func (customerService CustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return customerService.repo.FindAll()
}

func (customerService CustomerService) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return customerService.repo.FindCustomerById(id)
}

func (customerService CustomerService) GetCustomersByStatus(status bool) ([]domain.Customer, *errs.AppError) {
	return customerService.repo.FindCustomersByStatus(status)
}

//constructor method to initiate CustomerService class/struct

func NewCustomerService(repository domain.ICustomerRepository) CustomerService {
	return CustomerService{repository}
}

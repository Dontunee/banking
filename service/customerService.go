package service

import "github.com/Dontunee/banking/domain"

type ICustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerById(id string) (*domain.Customer, error)
}

type CustomerService struct {
	repo domain.ICustomerRepository // inject customer repo interface
}

//method in Customer Service class/struct

func (customerService CustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return customerService.repo.FindAll()
}

func (customerService CustomerService) GetCustomerById(id string) (*domain.Customer, error) {
	return customerService.repo.FindCustomerById(id)
}

//constructor method to initiate CustomerService class/struct

func NewCustomerService(repository domain.ICustomerRepository) CustomerService {
	return CustomerService{repository}
}

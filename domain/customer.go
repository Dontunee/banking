package domain

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
	FindAll() ([]Customer, error)
	FindCustomerById(id string) (*Customer, error)
}

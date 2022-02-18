package domain

//stub implementation
//adapter

type CustomerRepositoryStub struct {
	customers []Customer
}

func (repoStub CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return repoStub.customers, nil
}

//helper function to create repository stub
//more like constructor method

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{1001, "Tunde", "lagos", "10001", "1996-04-03", true},
		{1002, "Femi", "lagos", "1005", "2000-04-03", false},
	}
	return CustomerRepositoryStub{customers: customers}
}

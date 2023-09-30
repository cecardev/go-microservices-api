package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "John", City: "New York", Zipcode: "123", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Jane", City: "New York", Zipcode: "123", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "1003", Name: "Bob", City: "New York", Zipcode: "123", DateOfBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}

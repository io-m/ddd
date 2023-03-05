package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
)

var (
	ErrorCustomerNotFound = errors.New("the customer is not found")
	ErrorCreateCustomer   = errors.New("the customer cannot be created")
	ErrorUpdateCustomer   = errors.New("the customer cannot be updated")
)

// ICustomerRepo is an interface that exposes set of function to interact with any repository that implements it
type ICustomerRepo interface {
	GetOne(id uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}

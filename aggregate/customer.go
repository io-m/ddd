// Package aggregate holds our aggregates that combines more entities into a full objects that we work on inside of business logic
package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/io-m/ddd/entity"
	"github.com/io-m/ddd/valueobject"
)

var (
	ErrorInvalidPerson = errors.New("Customer has to have a valid name")
)

// Customer combines item entity, transaction value object and person entity
// Person entity is a root entity which means that person ID is the main ID for Customer
type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is factory to create new Customer aggregate
// It validates few bad cases, such as if name is empty etc.
func NewCustomer(name string) (Customer, error) {
	// basic validation for empty name passed
	if name == "" {
		return Customer{}, ErrorInvalidPerson
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

/* Getters and setters for private Customer fields */

func (c *Customer) GetCustomerId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetCustomerId(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetCustomerName() string {
	return c.person.Name
}

func (c *Customer) SetCustomerName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetCustomerAg() int {
	return c.person.Age
}

func (c *Customer) SetCustomerAge(age int) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Age = age
}

// Package memo is an implementation of ICustomerRepo interface to store Customers in-memory
package memo

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
	"github.com/io-m/ddd/domain/customer"
)

type MemoRepo struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoRepo {
	return &MemoRepo{
		customers: make(map[uuid.UUID]aggregate.Customer, 0),
	}
}

func (mr *MemoRepo) GetOne(id uuid.UUID) (aggregate.Customer, error) {
	c, ok := mr.customers[id]
	if !ok {
		return aggregate.Customer{}, customer.ErrorCustomerNotFound
	}
	return c, nil
}
func (mr *MemoRepo) Add(c aggregate.Customer) error {
	// Factory should protect us with creating map, but for extra safety, we create it if it is nil
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	id := c.GetCustomerId()
	// Make sure customer is not in repo
	if _, ok := mr.customers[id]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrorCreateCustomer)
	}
	mr.Lock()
	mr.customers[id] = c
	mr.Unlock()
	return nil
}
func (mr *MemoRepo) Update(c aggregate.Customer) error {
	id := c.GetCustomerId()
	// Create one if it does not exist
	if _, ok := mr.customers[id]; !ok {
		log.Printf("Customer %v does not exist while trying to update, hence creating one", c)
		mr.Add(c)
	}
	mr.Lock()
	mr.customers[id] = c
	mr.Unlock()
	return nil
}

package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/ramses2099/dddingolang/aggregate"
	"github.com/ramses2099/dddingolang/domain/customer"
)

//MemoryRepository fulfills the CustomerRepository interface
type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

//New is a factory function to generate a new of customers
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get finds a customer by ID
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// Add will new customer to the repository
func (mr *MemoryRepository) Add(aggregate.Customer) error {
	return nil
}

// Update will replace an exiting customer information with the new customer information
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if mr.customers == nil {
		//Safty check if customers is not create, shouldn't happen if usin the Factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	//Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists %w", customer.ErrFailedToAddCustomer)
	}
	//
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

package services

import (
	"github.com/google/uuid"
	"github.com/ramses2099/dddingolang/domain/customer"
)

// OrderConfiguration is on alias for a fuction that will takte in a pointer on OrderService and modify it
type OrderConfiguration func(os *OrderService) error

//OrderConfiguration is an implementation of the Orderservice
type OrderService struct {
	customers customer.CustomerRepository
}

// NewOrderService takes a varible amount of OrderConfiguration functions and returns a new OrderService
// Each Orderconfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the orderservice
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil

}

//CreateOrder will chaintogther all respository to create a order for customer
func (o *OrderService) CrateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	//Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}

	return nil
}

package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ramses2099/dddingolang/entity"
	"github.com/ramses2099/dddingolang/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

//Customer is a aggregate that comines all entities needed to represnt a customer
type Customer struct {
	//person is the root entity of a customer
	//wich means the person.ID is the main identifier for this aggregate
	person *entity.Person

	// a customer can hold many products
	produts []*entity.Item

	//a customer can perform many transaction
	transactions []valueobject.Transaction
}

//NewCustomer is a factory to create a new Customer aggregate
//It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// valid the name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Create a new person and generate ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	// Create a customer and initialize all the values to avoid nil pointer exceptions
	return Customer{
		person:       person,
		produts:      make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil

}

//GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

//SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if name == "" {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

//GetName
func (c *Customer) GetName() string {
	return c.person.Name
}

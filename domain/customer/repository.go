package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ramses2099/dddingolang/aggregate"
)

var (
	//ErrCustomerNotFound is returned when a customer is not found
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	//ErrFailedToAddCustomer is returned when the customer could not be added to the repository
	ErrFailedToAddCustomer = errors.New("failded to add the customer to the repository")
	//ErrUpdateCustomer is returned when the customer could not be updated in the repository
	ErrUpdateCustomer = errors.New("failed to update customer in the repository")
)

// CustomerRepository is a interface that defines the rules arround what a customer repository
// Has to be able to perform
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}

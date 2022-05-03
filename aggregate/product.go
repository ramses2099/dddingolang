package aggregate

import (
	"errors"

	"github.com/ramses2099/dddingolang/entity"
)

var (
	// ErrMissingValues is returned when a product is created without a name or decription
	ErrMissingValues = errors.New("missing values")
)

// Product is a aggregate that combines item with a price and quantity
type Product struct {
	// item is the root entity which is an item
	item  *entity.Item
	price float64
	// quantity is the number of product in stock
	quantiy int
}

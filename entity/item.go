package entity

import "github.com/google/uuid"

// Item represnts a Item for all sub dommains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}

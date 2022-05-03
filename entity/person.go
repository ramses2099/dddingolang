package entity

import "github.com/google/uuid"

// Person is an entity that representa a person in all Domains
type Person struct {
	// ID is the identifier of the entity, the ID is shared for all sub domains
	ID   uuid.UUID
	Name string
	Age  int
}

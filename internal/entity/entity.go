package entity

import uuid "github.com/satori/go.uuid"

// ID entity ID
type ID = uuid.UUID

// NewId creates a new ID entity
func NewID() ID {
	return ID(uuid.NewV4())
}

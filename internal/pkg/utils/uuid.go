package utils

import (
	"github.com/google/uuid"
)

func CreateID() uuid.UUID {
	return uuid.New()
}

func ParseUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

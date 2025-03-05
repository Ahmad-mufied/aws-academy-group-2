package utils

import (
	"fmt"
	"github.com/google/uuid"
)

// StringToUUID validates whether a string is a valid UUID v4.
// Returns the parsed uuid.UUID if valid, or an error if invalid.
func StringToUUID(s string) (uuid.UUID, error) {
	// Attempt to parse the string as a UUID
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, fmt.Errorf("string is not a valid UUID: %w", err)
	}

	// Check if the UUID is version 4
	if id.Version() != 4 {
		return uuid.Nil, fmt.Errorf("UUID is not version 4, but version %d", id.Version())
	}

	// Return the valid UUID v4
	return id, nil
}

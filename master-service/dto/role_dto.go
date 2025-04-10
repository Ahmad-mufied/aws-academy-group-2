package dto

import "github.com/google/uuid"

type CreateOrUpdateRoleRequest struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	Name      string     `json:"name"`
	IsActive  *bool      `json:"is_active,omitempty"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
}

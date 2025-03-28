package model

import "time"

type AttributeField struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type StatusResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type AttributeResponse struct {
	Role   string `json:"role"`
	Status string `json:"status"`
}

type ProductResponse struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type UserResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	DoB       string            `json:"date_of_birth"`
	Role      RoleResponse      `json:"role,omitempty"`
	Status    StatusResponse    `json:"status,omitempty"`
	Product   []ProductResponse `json:"product,omitempty"`
	CretedAt  string            `json:"created_at,omitempty"`
	UpdatedAt string            `json:"updated_at,omitempty"`
}

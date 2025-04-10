package dto

import "time"

type Product struct {
	ProductID string    `json:"product_id"`
	Name      string    `json:"name"`
	IsActive  string    `json:"is_active"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

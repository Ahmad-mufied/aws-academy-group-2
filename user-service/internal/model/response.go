package model

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
	Attribute AttributeResponse `json:"attribute"`
	Product   []ProductResponse `json:"product,omitempty"`
	CretedAt  string            `json:"created_at,omitempty"`
	UpdatedAt string            `json:"updated_at,omitempty"`
}

package model

type AttributeResponse struct {
	DoB    string `json:"date_of_birth"`
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
	Attribute AttributeResponse `json:"attribute"`
}

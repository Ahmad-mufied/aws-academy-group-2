package swagger

type CreateUserRequest struct {
	Name     string `json:"name" example:"David Afdal"`
	Email    string `json:"email" example:"david.afdal@example.com"`
	DoB      string `json:"date_of_birth" example:"2000-01-01"`
	RoleID   string `json:"role_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StatusID string `json:"status_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" example:"David Afdal"`
	Email    string `json:"email" example:"david.afdal@example.com"`
	DoB      string `json:"date_of_birth" example:"2000-01-01"`
	RoleID   string `json:"role_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StatusID string `json:"status_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type AssignProductsRequest struct {
	ProductIds []string `json:"product_ids" example:"550e8400-e29b-41d4-a716-446655440000"`
}

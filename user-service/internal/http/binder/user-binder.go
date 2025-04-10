package binder

type GetUsersBinder struct {
	Search    string `query:"search"`
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	StartDate string `query:"start_date"`
	EndDate   string `query:"end_date"`
}

type UpdateUserBinder struct {
	ID       string `param:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	DoB      string `json:"date_of_birth" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
	StatusID string `json:"status_id" validate:"required"`
}

type CreateUserBinder struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	DoB      string `json:"date_of_birth" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
	StatusID string `json:"status_id" validate:"required"`
}

type AssignProductsBinder struct {
	UserID     string   `param:"id" validate:"required"`
	ProductIds []string `json:"product_ids" validate:"required"`
}

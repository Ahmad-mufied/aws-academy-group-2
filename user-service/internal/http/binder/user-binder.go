package binder

type GetUsersBinder struct {
	Search string `query:"search"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}

type CreateUserBinder struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	DoB      string `json:"date_of_birth" binding:"required"`
	RoleID   string `json:"role" binding:"required"`
	StatusID string `json:"status" binding:"required"`
}

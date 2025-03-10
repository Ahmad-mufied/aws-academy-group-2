package swagger

type Meta struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Success"`
}

type MetaBadRequest struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Invalid request"`
}

type MetaCreated struct {
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Successfully created user"`
}

type MetaConflict struct {
	Code    int    `json:"code" example:"409"`
	Message string `json:"message" example:"user already exists"`
}

type MetaNotFound struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"user not found"`
}

type MetaInternalServerError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal server error"`
}

type User struct {
	ID          string `json:"id" example:"e1107b61-3531-4834-be96-a2cf981bced9"`
	Name        string `json:"name" example:"David Afdal"`
	Email       string `json:"email" example:"david.afdal@example.com"`
	DateOfBirth string `json:"date_of_birth" example:"2000-01-01"`
	Attribute   struct {
		Role   string `json:"role" example:"admin"`
		Status string `json:"status" example:"active"`
	} `json:"attribute"`
	CreatedAt string `json:"created_at" example:"2025-03-10 17:31:48"`
	UpdatedAt string `json:"updated_at" example:"2025-03-10 17:31:48"`
}

type ValidationErrorData struct {
	DateOfBirth string `json:"date_of_birth" example:"DoB is required"`
	Email       string `json:"email" example:"Email must be a valid email address"`
	Name        string `json:"name" example:"Name is required"`
	RoleID      string `json:"role_id" example:"RoleID is required"`
	StatusID    string `json:"status_id" example:"StatusID is required"`
}

type UserListData struct {
	Limit      int    `json:"limit" example:"10"`
	Page       int    `json:"page" example:"1"`
	TotalData  int    `json:"total_data" example:"2"`
	TotalPages int    `json:"total_pages" example:"1"`
	Users      []User `json:"users"`
}

type UserListResponse struct {
	Meta Meta         `json:"meta"`
	Data UserListData `json:"data"`
}

type SingleUserResponse struct {
	Meta Meta `json:"meta"`
	Data User `json:"data"`
}

type CreatedResponse struct {
	Meta MetaCreated `json:"meta"`
	Data User        `json:"data"`
}

type DeletedResponse struct {
	Meta Meta    `json:"meta"`
	Data *string `json:"data" example:"null"`
}

type ValidationErrorResponse struct {
	Meta MetaBadRequest      `json:"meta"`
	Data ValidationErrorData `json:"data"`
}

// NotFoundResponse represents a 404 error response with no data
type NotFoundResponse struct {
	Meta MetaNotFound `json:"meta"`
	Data *string      `json:"data" example:"null"`
}

// InternalServerErrorResponse represents a 500 error response with no data
type InternalServerErrorResponse struct {
	Meta MetaInternalServerError `json:"meta"`
	Data *string                 `json:"data" example:"null"`
}

// ConflictResponse represents a 409 error response with a conflict message
type ConflictResponse struct {
	Meta MetaConflict `json:"meta"`
	Data *string      `json:"data" example:"null"`
}

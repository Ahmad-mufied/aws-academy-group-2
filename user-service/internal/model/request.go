package model

type CreateAttributeReqest struct {
	UserID   string `json:"user_id"`
	DoB      string `json:"date_of_birth"`
	RoleID   string `json:"role"`
	StatusID string `json:"status"`
}

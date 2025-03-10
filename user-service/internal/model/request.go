package model

type CreateAttributeReqest struct {
	UserID   string `json:"user_id"`
	RoleID   string `json:"role"`
	StatusID string `json:"status"`
}

package dto

type ProductResponse struct {
	Id        string `json:"product_id"`
	Name      string `json:"name"`
	IsActive  string `json:"is_active"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProductRequest struct {
	Name      string `json:"name" validate:"required"`
	IsActive  string `json:"is_active"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

func (p *ProductRequest) StatusAsBool() bool {
	return p.IsActive == "active"
}

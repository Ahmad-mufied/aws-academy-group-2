package domain

import (
	"context"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/dto"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID uuid.UUID `bson:"product_id"`
	Name      string    `bson:"name"`
	IsActive  bool      `bson:"is_active"`
	CreatedBy uuid.UUID `bson:"created_by"`
	UpdatedBy uuid.UUID `bson:"updated_by"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type ProductRepository interface {
	FindAll(context.Context) ([]*Product, error)
	FindByID(context.Context, uuid.UUID) (*Product, error)
	Create(context.Context, *Product) error
	Update(context.Context, *Product) error
	Delete(context.Context, uuid.UUID) error
	Exists(context.Context, uuid.UUID) (bool, error)
}

func (p *Product) ToDto() *dto.ProductResponse {
	return &dto.ProductResponse{
		Id:        p.ProductID.String(),
		Name:      p.Name,
		IsActive:  p.statusAsText(),
		CreatedBy: p.CreatedBy.String(),
		UpdatedBy: p.UpdatedBy.String(),
		CreatedAt: p.CreatedAt.Format(time.RFC3339),
		UpdatedAt: p.UpdatedAt.Format(time.RFC3339),
	}
}

func (p *Product) statusAsText() string {
	status := "active"
	if !p.IsActive {
		status = "inactive"
	}

	return status
}

package service

import (
	"context"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/domain"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/dto"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/logger"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/repository"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/utils"
	"github.com/google/uuid"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]*dto.ProductResponse, error)
	GetByID(ctx context.Context, productID string) (*dto.ProductResponse, error)
	Create(ctx context.Context, product *dto.ProductRequest) (*dto.ProductResponse, error)
	Update(ctx context.Context, productID string, product *dto.ProductRequest) (*dto.ProductResponse, error)
	Delete(ctx context.Context, productID string) error
}

type productService struct {
	repo repository.MongoProductRepository
}

func (p *productService) GetAll(ctx context.Context) ([]*dto.ProductResponse, error) {

	products, err := p.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	productResponses := make([]*dto.ProductResponse, 0)
	for _, product := range products {
		productResponses = append(productResponses, product.ToDto())
	}

	return productResponses, nil
}

func (p *productService) GetByID(ctx context.Context, productID string) (*dto.ProductResponse, error) {

	productUuid, err := utils.StringToUUID(productID)

	var product *domain.Product
	product, err = p.repo.FindByID(ctx, productUuid)
	if err != nil {
		return nil, err
	}

	return product.ToDto(), nil
}

func (p *productService) Create(ctx context.Context, product *dto.ProductRequest) (*dto.ProductResponse, error) {

	var createdBy uuid.UUID
	var updatedBy uuid.UUID

	var err error
	if product.CreatedBy != "" {
		createdBy, err = utils.StringToUUID(product.CreatedBy)
		if err != nil {
			return nil, err
		}
	}

	if product.UpdatedBy != "" {
		updatedBy, err = utils.StringToUUID(product.UpdatedBy)
		if err != nil {
			return nil, err
		}
	}

	newProduct := &domain.Product{
		Name:      product.Name,
		IsActive:  product.StatusAsBool(),
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}

	err = p.repo.Create(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	return newProduct.ToDto(), nil
}

func (p *productService) Update(ctx context.Context, productID string, product *dto.ProductRequest) (*dto.ProductResponse, error) {

	productUuid, err := utils.StringToUUID(productID)
	if err != nil {
		logger.Error("error converting product id to uuid", err)
		return nil, err
	}

	var createdBy uuid.UUID
	var updatedBy uuid.UUID

	if product.CreatedBy != "" {
		createdBy, err = utils.StringToUUID(product.CreatedBy)
		if err != nil {
			logger.Error("error converting created by to uuid", err)
			return nil, err
		}
	}

	if product.UpdatedBy != "" {
		updatedBy, err = utils.StringToUUID(product.UpdatedBy)
		if err != nil {
			logger.Error("error converting updated by to uuid", err)
			return nil, err
		}
	}

	productToUpdate := &domain.Product{
		ProductID: productUuid,
		Name:      product.Name,
		IsActive:  product.StatusAsBool(),
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}

	err = p.repo.Update(ctx, productToUpdate)
	if err != nil {
		return nil, err
	}

	return productToUpdate.ToDto(), nil
}

func (p *productService) Delete(ctx context.Context, productID string) error {

	productUuid, err := utils.StringToUUID(productID)
	if err != nil {
		return err
	}

	err = p.repo.Delete(ctx, productUuid)
	if err != nil {
		return err
	}

	return nil
}

func NewProductService(repo repository.MongoProductRepository) ProductService {
	return &productService{repo}
}

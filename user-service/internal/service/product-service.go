package service

import (
	"encoding/json"
	"fmt"

	"github.com/DavidAfdal/user-services/config"
	"github.com/DavidAfdal/user-services/internal/dto"
	"github.com/DavidAfdal/user-services/pkg/httpclient"
)

type ProductService interface {
	GetProducts() ([]dto.Product, error)
	FilterByIDs(ids []string, products []dto.Product) []dto.Product
	CheckProductExists(productIDs []string) bool
}
type productService struct {
	config     *config.Config
	httpClient *httpclient.HTTPClient
}

func NewProductService(httpClient *httpclient.HTTPClient, config *config.Config) ProductService {
	return &productService{httpClient: httpClient, config: config}
}

func (s *productService) GetProducts() ([]dto.Product, error) {
	products := make([]dto.Product, 0)

	resp, err := s.httpClient.Fetch(s.config.ProductApi)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(resp), &products); err != nil {
		fmt.Println("Error unmarshalling:", err)
		return nil, fmt.Errorf("failed to unmarshal products: %w", err)
	}
	return products, nil
}

func (s *productService) FilterByIDs(ids []string, products []dto.Product) []dto.Product {
	if len(ids) == 0 {
		return []dto.Product{}
	}
	var filtered []dto.Product
	idSet := make(map[string]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	for _, product := range products {
		if _, ok := idSet[product.ProductID]; ok {
			filtered = append(filtered, product)
		}
	}

	return filtered
}

func (s *productService) CheckProductExists(productIDs []string) bool {
	products, err := s.GetProducts()
	if err != nil {
		return false
	}

	productMap := make(map[string]bool)
	for _, product := range products {
		productMap[product.ProductID] = true
	}

	for _, productID := range productIDs {
		if !productMap[productID] {
			return false
		}
	}

	return true
}

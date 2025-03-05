package server

import (
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/dto"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	productResponses, err := h.productService.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productResponses)
}

func (h *ProductHandler) GetProductById(c echo.Context) error {
	id := c.Param("id")

	productResponse, err := h.productService.GetByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productResponse)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var productRequest dto.ProductRequest

	if err := c.Bind(&productRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	productResponse, err := h.productService.Create(c.Request().Context(), &productRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, productResponse)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("product_id")
	var productRequest dto.ProductRequest

	if err := c.Bind(&productRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	productResponse, err := h.productService.Update(c.Request().Context(), id, &productRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productResponse)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("product_id")

	err := h.productService.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

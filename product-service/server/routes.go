package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo, h *ProductHandler) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/products", h.GetAllProducts)
	e.GET("/products/:id", h.GetProductById)
	e.POST("/products", h.CreateProduct)
	e.PUT("/products/:product_id", h.UpdateProduct)
	e.DELETE("/products/:product_id", h.DeleteProduct)
}

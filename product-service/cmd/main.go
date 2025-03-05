package main

import (
	"context"
	"errors"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/config"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/repository"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/server"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	config.InitMongo()

	productRepository, _ := repository.NewMongoProductRepository(config.DB.Database("test-product-db"), nil)
	productService := service.NewProductService(*productRepository)
	productHandler := server.NewProductHandler(productService)

	// Start the server
	startAndGracefullyStopServer(echo.New(), productHandler)
}

func startAndGracefullyStopServer(e *echo.Echo, productHandler *server.ProductHandler) {

	// Add CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	// Register routes
	server.RegisterRoutes(e, productHandler)

	port := config.Viper.GetString("WEB_SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: e,
	}

	go func() {
		if err := e.StartServer(srv); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

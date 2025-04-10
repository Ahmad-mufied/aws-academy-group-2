package main

import (
	"github.com/DavidAfdal/user-services/config"
	"github.com/DavidAfdal/user-services/internal/builder"
	"github.com/DavidAfdal/user-services/pkg/database"
	"github.com/DavidAfdal/user-services/pkg/logger"
	"github.com/DavidAfdal/user-services/pkg/server"
)

// @title User Service API
// @version 1.0
// @description Documentation of Api for user services.

// @host localhost:8080
// @BasePath /api/v1
func main() {
	logger.InitLog()

	cfg, err := config.NewConfig(".env")
	checkError(err)

	database, err := database.InitDB(cfg.Database)
	checkError(err)

	publicRoutes := builder.BuildPublicRoutes(database, cfg)

	srv := server.NewServer(publicRoutes)

	srv.Run(cfg.Port)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/DavidAfdal/user-services/config"
	"github.com/DavidAfdal/user-services/internal/builder"
	"github.com/DavidAfdal/user-services/pkg/database"
	"github.com/DavidAfdal/user-services/pkg/server"
)

func main() {
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

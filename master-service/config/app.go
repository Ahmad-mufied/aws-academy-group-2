package config

import (
	"master-service/handler"
	"master-service/repository"
)

type App struct {
	RoleHandler   handler.RoleHandler
	StatusHandler handler.StatusHandler
}

func InitializeApp() *App {
	db := GetDB()

	// init repo
	roleRepo := repository.InitRoleRepository(db)
	statusRepo := repository.InitStatusRepository(db)

	// init handler
	roleHandler := handler.InitRoleHandler(roleRepo)
	statusHandler := handler.InitStatusHandler(statusRepo)

	return &App{
		RoleHandler:   roleHandler,
		StatusHandler: statusHandler,
	}
}

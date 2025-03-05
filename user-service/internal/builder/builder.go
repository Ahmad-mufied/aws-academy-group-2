package builder

import (
	"time"

	"github.com/DavidAfdal/user-services/config"
	"github.com/DavidAfdal/user-services/internal/http/handler"
	"github.com/DavidAfdal/user-services/internal/http/router"
	"github.com/DavidAfdal/user-services/internal/repository"
	"github.com/DavidAfdal/user-services/internal/service"
	"github.com/DavidAfdal/user-services/pkg/httpclient"
	"github.com/DavidAfdal/user-services/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB, cfg *config.Config) []*route.Route {

	httpClient := httpclient.NewHTTPClient(time.Second * 10)

	userRepo := repository.NewUserRepository(db)

	userServices := service.NewUserService(httpClient, userRepo, cfg)

	userHandler := handler.NewUserHandler(userServices)

	return router.PublicRoutes(userHandler)
}

package router

import (
	"github.com/DavidAfdal/user-services/internal/http/handler"
	"github.com/DavidAfdal/user-services/pkg/route"
)

func PublicRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  "GET",
			Path:    "/users",
			Handler: userHandler.GetUsers,
		},
	}
}

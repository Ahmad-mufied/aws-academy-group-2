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
		{
			Method:  "POST",
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},
		{
			Method:  "GET",
			Path:    "/users/:id",
			Handler: userHandler.GetUser,
		},
		{
			Method:  "PUT",
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
		},
		{
			Method:  "DELETE",
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
		},
	}
}

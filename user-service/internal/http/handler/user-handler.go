package handler

import (
	"net/http"

	"github.com/DavidAfdal/user-services/internal/http/binder"
	"github.com/DavidAfdal/user-services/internal/service"
	"github.com/DavidAfdal/user-services/pkg/validator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (r *UserHandler) GetUsers(c echo.Context) error {
	var input binder.GetUsersBinder

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errorMessage,
			"data":    data,
		})
	}

	users, total, err := r.userService.GetUsers(input)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
		"total": total,
	})
}

func checkValidation(input interface{}) (errorMessage string, data interface{}) {
	validationErrors := validator.Validate(input)
	if validationErrors != nil {
		if _, exists := validationErrors["error"]; exists {
			return "validasi input gagal", nil
		}
		return "validasi input gagal", validationErrors
	}
	return "", nil
}

package handler

import (
	"net/http"

	"github.com/DavidAfdal/user-services/internal/http/binder"
	"github.com/DavidAfdal/user-services/internal/service"
	"github.com/DavidAfdal/user-services/pkg/logger"
	"github.com/DavidAfdal/user-services/pkg/response"
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

// GetUsers godoc
// @Summary Get all users
// @Description Get paginated list of users
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of users per page"
// @Param search query string false "Search query for name or email"
// @Param start_date query string false "Start date format: YYYY-MM-DD"
// @Param end_date query string false "End date format: YYYY-MM-DD"
// @Success 200 {object} swagger.UserListResponse
// @Failure 500 {object} swagger.InternalServerErrorResponse
// @Router /users [get]
func (r *UserHandler) GetUsers(c echo.Context) error {
	var input binder.GetUsersBinder

	if err := c.Bind(&input); err != nil {
		logger.Error(c, "Failed to bind request in GetUsers")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		logger.Error(c, "Validation failed in GetUsers: "+errorMessage)
		return c.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	pagination, err := r.userService.GetUsers(input)

	if err != nil {
		logger.Error(c, err.Error())
		return c.JSON(err.StatusCode, response.ErrorResponse(err.StatusCode, err.Message))
	}

	logger.Info(c, "Successfully fetched users")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"meta": map[string]interface{}{
			"page":        pagination.Page,
			"limit":       pagination.PageSize,
			"total_data":  pagination.TotalCount,
			"total_pages": pagination.TotalPages,
		},
		"data": pagination.Data,
	})
}

// CreateUser godoc
// @Summary Create a new user
// @Description Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body swagger.CreateUserRequest true "User data"
// @Success 201 {object} swagger.CreatedResponse
// @Failure 400 {object} swagger.ValidationErrorResponse
// @Failure 409 {object} swagger.ConflictResponse
// @Failure 500 {object} swagger.InternalServerErrorResponse
// @Router /users [post]
func (r *UserHandler) CreateUser(c echo.Context) error {
	var input binder.CreateUserBinder

	if err := c.Bind(&input); err != nil {
		logger.Error(c, "Failed to bind request in CreateUser")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		logger.Warn(c, "Validation failed in CreateUser: "+errorMessage)
		return c.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	user, err := r.userService.CreateUser(input)
	if err != nil {
		logger.Error(c, err.Error())
		return c.JSON(err.StatusCode, response.ErrorResponse(err.StatusCode, err.Message))
	}

	logger.Info(c, "Successfully created user in CreateUser")
	return c.JSON(http.StatusCreated, response.SuccessResponse(http.StatusCreated, "Successfully created user", user))
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get user details
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} swagger.SingleUserResponse
// @Failure 404 {object} swagger.NotFoundResponse
// @Failure 500 {object} swagger.InternalServerErrorResponse
// @Router /users/{id} [get]
func (r *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := r.userService.GetUser(id)

	if err != nil {
		logger.Error(c, err.Error())
		return c.JSON(err.StatusCode, response.ErrorResponse(err.StatusCode, err.Message))
	}

	logger.Info(c, "Successfully fetched user in GetUser")

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully fetched user", user))
}

// UpdateUser godoc
// @Summary Update user details
// @Description Update existing user information
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body swagger.UpdateUserRequest true "User data"
// @Success 200 {object} swagger.SingleUserResponse
// @Failure 400 {object} swagger.ValidationErrorResponse
// @Failure 404 {object} swagger.NotFoundResponse
// @Failure 500 {object} swagger.InternalServerErrorResponse
// @Router /users/{id} [put]
func (r *UserHandler) UpdateUser(c echo.Context) error {
	var input binder.UpdateUserBinder

	if err := c.Bind(&input); err != nil {
		logger.Error(c, "Failed to bind request in UpdateUser")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		logger.Warn(c, "Validation failed in UpdateUser: "+errorMessage)
		return c.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	user, err := r.userService.UpdateUser(input)
	if err != nil {
		logger.Error(c, err.Error())
		return c.JSON(err.StatusCode, response.ErrorResponse(err.StatusCode, err.Message))
	}

	logger.Info(c, "Successfully updated user in UpdateUser")
	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully updated user", user))
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Remove user from database
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} swagger.DeletedResponse
// @Failure 404 {object} swagger.NotFoundResponse
// @Failure 500 {object} swagger.InternalServerErrorResponse
// @Router /users/{id} [delete]
func (r *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if err := r.userService.DeleteUser(id); err != nil {
		logger.Error(c, err.Error())
		return c.JSON(err.StatusCode, response.ErrorResponse(err.StatusCode, err.Message))
	}

	logger.Info(c, "Successfully deleted user in DeleteUser")
	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully deleted user", nil))
}

// AssignProducts godoc
// @Summary Assign products to user
// @Description Assign products to user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} swagger.NotFoundResponse
// @Failure 500 {object} swagger.InternalServerErrorResponse
// @Router /users/{id}/assign-products [post]
func (r *UserHandler) AssignProducts(c echo.Context) error {
	var input binder.AssignProductsBinder

	if err := c.Bind(&input); err != nil {
		logger.Error(c, "Failed to bind request in AssignProducts")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		logger.Warn(c, "Validation failed in AssignProducts: "+errorMessage)
		return c.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	if err := r.userService.AssignProducts(input); err != nil {
		logger.Error(c, err.Error())
		return c.JSON(err.StatusCode, response.ErrorResponse(err.StatusCode, err.Message))
	}

	logger.Info(c, "Successfully assigned products to user in AssignProducts")
	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully assigned products to user", nil))
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

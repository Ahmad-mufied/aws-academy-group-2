package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DavidAfdal/user-services/config"
	"github.com/DavidAfdal/user-services/internal/dto"
	"github.com/DavidAfdal/user-services/internal/http/binder"
	"github.com/DavidAfdal/user-services/internal/model"
	"github.com/DavidAfdal/user-services/internal/repository"
	"github.com/DavidAfdal/user-services/pkg/constant"
	exceptions "github.com/DavidAfdal/user-services/pkg/execptions"
	"github.com/DavidAfdal/user-services/pkg/pagination"
	"github.com/google/uuid"
)

type UserService interface {
	GetUsers(input binder.GetUsersBinder) (pagination.Pagination, *exceptions.HTTPError)
	CreateUser(input binder.CreateUserBinder) (dto.UserResponse, *exceptions.HTTPError)
	GetUser(id string) (dto.UserResponse, *exceptions.HTTPError)
	UpdateUser(input binder.UpdateUserBinder) (dto.UserResponse, *exceptions.HTTPError)
	DeleteUser(id string) *exceptions.HTTPError
	AssignProducts(input binder.AssignProductsBinder) *exceptions.HTTPError
}

type userService struct {
	config         *config.Config
	userRepo       repository.UserRepository
	productService ProductService
	masterService  MasterService
}

func NewUserService(userRepo repository.UserRepository, config *config.Config, productService ProductService, masterService MasterService) UserService {
	return &userService{userRepo: userRepo, config: config, productService: productService, masterService: masterService}
}

func (s *userService) GetUsers(input binder.GetUsersBinder) (pagination.Pagination, *exceptions.HTTPError) {

	if input.Limit == 0 {
		input.Limit = 10
	}

	if input.Page == 0 {
		input.Page = 1
	}

	users, total, err := s.userRepo.GetUsers(input.Search, input.Page, input.Limit, input.StartDate, input.EndDate)

	if err != nil {
		return pagination.Pagination{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	usersResponse, err := s.usersWithAttributes(users)

	if err != nil {
		return pagination.Pagination{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return pagination.Paginate(usersResponse, total, input.Page, input.Limit), nil
}

func (s *userService) GetUser(id string) (dto.UserResponse, *exceptions.HTTPError) {

	user, err := s.userRepo.GetUser(id)

	if err != nil {
		if err == constant.ErrUserNotFound {
			return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusNotFound, constant.ErrUserNotFound.Error())
		}
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userResponse, err := s.userWithAttributes(*user)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return userResponse, nil
}

func (s *userService) CreateUser(input binder.CreateUserBinder) (dto.UserResponse, *exceptions.HTTPError) {

	role, err := s.masterService.GetRoleByID(input.RoleID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusNotFound, "Role not found")
	}

	status, err := s.masterService.GetStatusByID(input.StatusID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusNotFound, "Status not found")
	}

	roleID, err := uuid.Parse(role.ID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	statusID, err := uuid.Parse(status.ID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	DoB, _ := time.Parse("2006-01-02", input.DoB)

	user := model.User{Name: input.Name, Email: input.Email, RoleID: roleID, StatusID: statusID, DateOfBirth: DoB}

	createdUser, err := s.userRepo.CreateUser(&user)

	if err != nil {
		if err == constant.ErrUserExists {
			return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusConflict, constant.ErrUserExists.Error())
		}
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return createdUser.ToDto(dto.AttributeResponse{
		Role:   *role,
		Status: *status,
	}, []dto.Product{}), nil
}

func (s *userService) UpdateUser(input binder.UpdateUserBinder) (dto.UserResponse, *exceptions.HTTPError) {

	_, err := s.userRepo.GetUser(input.ID)

	if err != nil && err == constant.ErrUserNotFound {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusNotFound, constant.ErrUserNotFound.Error())
	}

	role, err := s.masterService.GetRoleByID(input.RoleID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	status, err := s.masterService.GetStatusByID(input.StatusID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	DoB, _ := time.Parse("2006-01-02", input.DoB)

	id, err := uuid.Parse(input.ID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	roleID, err := uuid.Parse(role.ID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	statusID, err := uuid.Parse(status.ID)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user := &model.User{ID: id, RoleID: roleID, StatusID: statusID, Name: input.Name, Email: input.Email, DateOfBirth: DoB}

	updatedUser, err := s.userRepo.UpdateUser(user)

	if err != nil {
		return dto.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return updatedUser.ToDto(dto.AttributeResponse{}, []dto.Product{}), nil
}

func (s *userService) DeleteUser(id string) *exceptions.HTTPError {

	_, err := s.userRepo.GetUser(id)

	if err != nil && err == constant.ErrUserNotFound {
		return exceptions.NewHTTPError(http.StatusNotFound, constant.ErrUserNotFound.Error())
	}

	if err := s.userRepo.DeleteUser(id); err != nil {
		return exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s *userService) AssignProducts(input binder.AssignProductsBinder) *exceptions.HTTPError {

	user, err := s.userRepo.GetUser(input.UserID)

	if err != nil && err == constant.ErrUserNotFound {
		return exceptions.NewHTTPError(http.StatusNotFound, constant.ErrUserNotFound.Error())
	}

	if productExits := s.productService.CheckProductExists(input.ProductIds); !productExits {
		return exceptions.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	if err := s.userRepo.AssignProducts(user, input.ProductIds); err != nil {
		return exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s *userService) usersWithAttributes(users []model.User) ([]dto.UserResponse, error) {
	userResponse := make([]dto.UserResponse, 0)

	if len(users) == 0 {
		return userResponse, nil
	}

	roles, errRoles := s.masterService.GetRoles()
	status, errProducts := s.masterService.GetStatus()

	if errRoles != nil || errProducts != nil {
		return nil, errors.New("internal server error")
	}

	for _, user := range users {
		var productIDs []string

		if len(user.ProductIDs) == 0 || string(user.ProductIDs) == "null" {
			productIDs = []string{}
		} else {
			err := json.Unmarshal(user.ProductIDs, &productIDs)
			if err != nil {
				fmt.Println("Error unmarshalling:", err)
				return nil, err
			}
		}

		role := s.masterService.FilterRoleByID(user.RoleID.String(), roles)

		statusData := s.masterService.FilterStatusByID(user.StatusID.String(), status)

		response := user.ToDto(dto.AttributeResponse{
			Role:   role,
			Status: statusData,
		}, []dto.Product{})

		userResponse = append(userResponse, response)

	}

	return userResponse, nil
}

func (s *userService) userWithAttributes(user model.User) (dto.UserResponse, error) {

	products, err := s.productService.GetProducts()

	if err != nil {
		return dto.UserResponse{}, err
	}

	var productIDs []string

	if len(user.ProductIDs) == 0 || string(user.ProductIDs) == "null" {
		productIDs = []string{}
	} else {
		err := json.Unmarshal(user.ProductIDs, &productIDs)
		if err != nil {
			fmt.Println("Error unmarshalling:", err)
			return dto.UserResponse{}, err
		}
	}

	productData := s.productService.FilterByIDs(productIDs, products)

	role, err := s.masterService.GetRoleByID(user.RoleID.String())

	if err != nil {
		return dto.UserResponse{}, err
	}

	status, err := s.masterService.GetStatusByID(user.StatusID.String())

	if err != nil {
		return dto.UserResponse{}, err
	}

	response := user.ToDto(dto.AttributeResponse{
		Role:   *role,
		Status: *status,
	}, productData)

	return response, nil
}

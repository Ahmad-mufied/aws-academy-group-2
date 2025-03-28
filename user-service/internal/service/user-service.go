package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DavidAfdal/user-services/config"
	"github.com/DavidAfdal/user-services/internal/http/binder"
	"github.com/DavidAfdal/user-services/internal/model"
	"github.com/DavidAfdal/user-services/internal/repository"
	"github.com/DavidAfdal/user-services/pkg/constant"
	exceptions "github.com/DavidAfdal/user-services/pkg/execptions"
	"github.com/DavidAfdal/user-services/pkg/httpclient"
	"github.com/DavidAfdal/user-services/pkg/pagination"
	"github.com/google/uuid"
)

type UserService interface {
	GetUsers(input binder.GetUsersBinder) (pagination.Pagination, *exceptions.HTTPError)
	CreateUser(input binder.CreateUserBinder) (model.UserResponse, *exceptions.HTTPError)
	GetUser(id string) (model.UserResponse, *exceptions.HTTPError)
	UpdateUser(input binder.UpdateUserBinder) (model.UserResponse, *exceptions.HTTPError)
	DeleteUser(id string) *exceptions.HTTPError
}

type userService struct {
	httpClient *httpclient.HTTPClient
	userRepo   repository.UserRepository
	config     *config.Config
}

func NewUserService(httpClient *httpclient.HTTPClient, userRepo repository.UserRepository, config *config.Config) UserService {
	return &userService{httpClient: httpClient, userRepo: userRepo, config: config}
}

func (s *userService) GetUsers(input binder.GetUsersBinder) (pagination.Pagination, *exceptions.HTTPError) {
	usersResponse := make([]model.UserResponse, 0)

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

	for _, user := range users {

		role, err := s.getRole(user.RoleID.String())

		if err != nil {
			return pagination.Pagination{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		status, err := s.getStatus(user.StatusID.String())

		if err != nil {
			return pagination.Pagination{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		usersResponse = append(usersResponse, model.UserResponse{ID: user.ID.String(), Name: user.Name, DoB: user.DateOfBirth.Format("2006-01-02"), Email: user.Email, Role: role, Status: status, CretedAt: user.CreatedAt.Format("2006-01-02 15:04:05"), UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05")})
	}

	return pagination.Paginate(usersResponse, total, input.Page, input.Limit), nil
}

func (s *userService) GetUser(id string) (model.UserResponse, *exceptions.HTTPError) {

	user, err := s.userRepo.GetUser(id)

	if err != nil {
		if err == constant.ErrUserNotFound {
			return model.UserResponse{}, exceptions.NewHTTPError(http.StatusNotFound, constant.ErrUserNotFound.Error())
		}
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	role, err := s.getRole(user.RoleID.String())

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	status, err := s.getStatus(user.StatusID.String())

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return model.UserResponse{ID: user.ID.String(), Name: user.Name, Email: user.Email, DoB: user.DateOfBirth.Format("2006-01-02"), CretedAt: user.CreatedAt.Format("2006-01-02 15:04:05"), UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"), Role: role, Status: status}, nil
}

func (s *userService) CreateUser(input binder.CreateUserBinder) (model.UserResponse, *exceptions.HTTPError) {

	role, err := s.getRole(input.RoleID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	status, err := s.getStatus(input.StatusID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	roleID, err := uuid.Parse(input.RoleID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	statusID, err := uuid.Parse(input.StatusID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	DoB, _ := time.Parse("2006-01-02", input.DoB)

	user := model.User{Name: input.Name, Email: input.Email, RoleID: roleID, StatusID: statusID, DateOfBirth: DoB}

	createdUser, err := s.userRepo.CreateUser(&user)

	if err != nil {
		if err == constant.ErrUserExists {
			return model.UserResponse{}, exceptions.NewHTTPError(http.StatusConflict, constant.ErrUserExists.Error())
		}
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return model.UserResponse{ID: createdUser.ID.String(), Name: createdUser.Name, Email: createdUser.Email, DoB: createdUser.DateOfBirth.Format("2006-01-02"), Role: role, Status: status}, nil
}

func (s *userService) UpdateUser(input binder.UpdateUserBinder) (model.UserResponse, *exceptions.HTTPError) {

	_, err := s.userRepo.GetUser(input.ID)

	if err != nil && err == constant.ErrUserNotFound {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusNotFound, constant.ErrUserNotFound.Error())
	}

	role, err := s.getRole(input.RoleID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	status, err := s.getStatus(input.StatusID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	DoB, _ := time.Parse("2006-01-02", input.DoB)

	id, err := uuid.Parse(input.ID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	roleID, err := uuid.Parse(input.RoleID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	statusID, err := uuid.Parse(input.StatusID)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user := &model.User{ID: id, RoleID: roleID, StatusID: statusID, Name: input.Name, Email: input.Email, DateOfBirth: DoB}

	updatedUser, err := s.userRepo.UpdateUser(user)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return model.UserResponse{ID: updatedUser.ID.String(), Name: updatedUser.Name, Email: updatedUser.Email, DoB: updatedUser.DateOfBirth.Format("2006-01-02"), Role: role, Status: status}, nil
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

func (s *userService) getRole(id string) (model.RoleResponse, error) {
	var role model.AttributeField

	url := fmt.Sprintf("%s/%s", s.config.MasterApi, id)

	resp, err := s.httpClient.Fetch(url)

	if err != nil {
		return model.RoleResponse{}, err
	}

	if err := json.Unmarshal([]byte(resp), &role); err != nil {
		return model.RoleResponse{}, err
	}

	return model.RoleResponse{ID: role.ID, Name: role.Name, IsActive: role.IsActive}, nil
}

func (s *userService) getStatus(id string) (model.StatusResponse, error) {
	var status model.AttributeField

	url := fmt.Sprintf("%s/%s", s.config.MasterApi, id)

	resp, err := s.httpClient.Fetch(url)

	if err != nil {
		return model.StatusResponse{}, err
	}

	if err := json.Unmarshal([]byte(resp), &status); err != nil {
		return model.StatusResponse{}, err
	}

	return model.StatusResponse{ID: status.ID, Name: status.Name, IsActive: status.IsActive}, nil
}

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
		var attributes model.AttributeResponse

		// attributes, err = s.getAttributeUser(user.ID.String())

		// if err != nil {
		// 	return pagination.Pagination{}, err
		// }

		usersResponse = append(usersResponse, model.UserResponse{ID: user.ID.String(), Name: user.Name, DoB: user.DateOfBirth.Format("2006-01-02"), Email: user.Email, Attribute: attributes, CretedAt: user.CreatedAt.Format("2006-01-02 15:04:05"), UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05")})
	}

	return pagination.Paginate(usersResponse, total, input.Page, input.Limit), nil
}

func (s *userService) GetUser(id string) (model.UserResponse, *exceptions.HTTPError) {
	var attributes model.AttributeResponse

	user, err := s.userRepo.GetUser(id)

	if err != nil {
		if err == constant.ErrUserNotFound {
			return model.UserResponse{}, exceptions.NewHTTPError(http.StatusNotFound, constant.ErrUserNotFound.Error())
		}
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// attributes, err = s.getAttributeUser(user.ID.String())

	// if err != nil {
	// 	return model.UserResponse{}, err
	// }

	return model.UserResponse{ID: user.ID.String(), Name: user.Name, Email: user.Email, DoB: user.DateOfBirth.Format("2006-01-02"), CretedAt: user.CreatedAt.Format("2006-01-02 15:04:05"), UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"), Attribute: attributes}, nil
}

func (s *userService) CreateUser(input binder.CreateUserBinder) (model.UserResponse, *exceptions.HTTPError) {

	DoB, _ := time.Parse("2006-01-02", input.DoB)

	user := model.User{Name: input.Name, Email: input.Email, DateOfBirth: DoB}

	createdUser, err := s.userRepo.CreateUser(&user)

	if err != nil {
		if err == constant.ErrUserExists {
			return model.UserResponse{}, exceptions.NewHTTPError(http.StatusConflict, constant.ErrUserExists.Error())
		}
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return model.UserResponse{ID: createdUser.ID.String(), Name: createdUser.Name, Email: createdUser.Email, DoB: createdUser.DateOfBirth.Format("2006-01-02")}, nil
}

func (s *userService) UpdateUser(input binder.UpdateUserBinder) (model.UserResponse, *exceptions.HTTPError) {

	_, err := s.userRepo.GetUser(input.ID)

	if err != nil && err == constant.ErrUserNotFound {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusNotFound, constant.ErrUserNotFound.Error())
	}

	var attributes model.AttributeResponse
	DoB, _ := time.Parse("2006-01-02", input.DoB)

	user := &model.User{Name: input.Name, Email: input.Email, DateOfBirth: DoB}

	updatedUser, err := s.userRepo.UpdateUser(user)

	if err != nil {
		return model.UserResponse{}, exceptions.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return model.UserResponse{ID: updatedUser.ID.String(), Name: updatedUser.Name, Email: updatedUser.Email, DoB: updatedUser.DateOfBirth.Format("2006-01-02"), Attribute: attributes}, nil
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

func (s *userService) getAttributeUser(id string) (model.AttributeResponse, error) {
	var attributes model.AttributeResponse

	url := fmt.Sprintf("%s/%s", s.config.MasterApi, id)

	resp, err := s.httpClient.Fetch(url)

	if err != nil {
		return model.AttributeResponse{}, err
	}

	if err := json.Unmarshal([]byte(resp), &attributes); err != nil {
		return model.AttributeResponse{}, err
	}

	return attributes, nil
}

func (s *userService) assignRoleAndStatus(userID, roleID, statusID string) error {
	var attributes model.AttributeResponse

	bodyRequest := model.CreateAttributeReqest{UserID: userID, RoleID: roleID, StatusID: statusID}

	resp, err := s.httpClient.Post(s.config.MasterApi, bodyRequest)

	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(resp), &attributes); err != nil {
		return err
	}

	return nil
}

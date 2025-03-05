package service

import (
	"encoding/json"
	"fmt"

	"github.com/DavidAfdal/user-services/config"
	"github.com/DavidAfdal/user-services/internal/http/binder"
	"github.com/DavidAfdal/user-services/internal/model"
	"github.com/DavidAfdal/user-services/internal/repository"
	"github.com/DavidAfdal/user-services/pkg/httpclient"
)

type UserService interface {
	GetUsers(input binder.GetUsersBinder) ([]model.UserResponse, int, error)
	CreateUser(input binder.CreateUserBinder) (model.UserResponse, error)
}

type userService struct {
	httpClient *httpclient.HTTPClient
	userRepo   repository.UserRepository
	config     *config.Config
}

func NewUserService(httpClient *httpclient.HTTPClient, userRepo repository.UserRepository, config *config.Config) UserService {
	return &userService{httpClient: httpClient, userRepo: userRepo, config: config}
}

func (s *userService) GetUsers(input binder.GetUsersBinder) ([]model.UserResponse, int, error) {
	usersResponse := make([]model.UserResponse, 0)
	users, total, err := s.userRepo.GetUsers(input.Search, input.Page, input.Limit)

	if err != nil {
		return nil, 0, err
	}

	for _, user := range users {
		var attributes model.AttributeResponse

		resp, err := s.httpClient.Fetch(fmt.Sprintf("%s/%s", s.config.MasterApi, user.ID.String()))

		if err != nil {
			return nil, 0, err
		}

		if err := json.Unmarshal([]byte(resp), &attributes); err != nil {
			return nil, 0, err
		}

		usersResponse = append(usersResponse, model.UserResponse{ID: user.ID.String(), Name: user.Name, Email: user.Email, Attribute: attributes})
	}

	return usersResponse, total, nil
}

func (s *userService) CreateUser(input binder.CreateUserBinder) (model.UserResponse, error) {
	var attributes model.AttributeResponse
	user := model.User{Name: input.Name, Email: input.Email}

	createdUser, err := s.userRepo.CreateUser(&user)

	if err != nil {
		return model.UserResponse{}, err
	}

	bodyRequest := model.CreateAttributeReqest{UserID: createdUser.ID.String(), DoB: input.DoB, RoleID: input.RoleID, StatusID: input.StatusID}

	resp, err := s.httpClient.Post(s.config.MasterApi, bodyRequest)

	if err != nil {
		return model.UserResponse{}, err
	}

	if err := json.Unmarshal([]byte(resp), &attributes); err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{ID: createdUser.ID.String(), Name: createdUser.Name, Email: createdUser.Email, Attribute: attributes}, nil
}

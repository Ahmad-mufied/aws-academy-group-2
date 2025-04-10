package service

import (
	"encoding/json"
	"fmt"

	"github.com/DavidAfdal/user-services/config"
	"github.com/DavidAfdal/user-services/internal/dto"
	"github.com/DavidAfdal/user-services/pkg/httpclient"
)

type MasterService interface {
	GetRoles() ([]dto.Role, error)
	GetRoleByID(id string) (*dto.Role, error)
	GetStatusByID(id string) (*dto.Status, error)
	GetStatus() ([]dto.Status, error)
	FilterRoleByID(id string, roles []dto.Role) dto.Role
	FilterStatusByID(id string, status []dto.Status) dto.Status
}

type masterService struct {
	httpClient *httpclient.HTTPClient
	config     *config.Config
}

func NewMasterService(httpClient *httpclient.HTTPClient, config *config.Config) MasterService {
	return &masterService{httpClient: httpClient, config: config}
}

func (s *masterService) GetRoles() ([]dto.Role, error) {
	resp, err := s.httpClient.Fetch(s.config.MasterApi + "/roles")

	if err != nil {
		return nil, err
	}

	var roles []dto.Role

	if err := json.Unmarshal([]byte(resp), &roles); err != nil {

		return nil, err
	}

	return roles, nil
}

func (s *masterService) GetRoleByID(id string) (*dto.Role, error) {
	url := fmt.Sprintf("%s/roles/%s", s.config.MasterApi, id)

	resp, err := s.httpClient.Fetch(url)

	if err != nil {
		return nil, err
	}

	var roles dto.Role

	if err := json.Unmarshal([]byte(resp), &roles); err != nil {
		fmt.Println("Error unmarshalling Role:", err)
		return nil, err
	}

	return &roles, nil
}

func (s *masterService) GetStatusByID(id string) (*dto.Status, error) {
	url := fmt.Sprintf("%s/status/%s", s.config.MasterApi, id)
	resp, err := s.httpClient.Fetch(url)

	if err != nil {
		return nil, err
	}

	var status dto.Status

	if err := json.Unmarshal([]byte(resp), &status); err != nil {
		fmt.Println("Error unmarshalling Status:", err)
		return nil, err
	}

	return &status, nil
}

func (s *masterService) GetStatus() ([]dto.Status, error) {
	resp, err := s.httpClient.Fetch(s.config.MasterApi + "/status")

	if err != nil {
		return nil, err
	}

	var status []dto.Status

	if err := json.Unmarshal([]byte(resp), &status); err != nil {
		return nil, err
	}

	return status, nil
}

func (s *masterService) FilterRoleByID(id string, roles []dto.Role) dto.Role {
	for _, role := range roles {
		if role.ID == id {
			fmt.Println("Role found:", role)
			return role
		}
	}

	return dto.Role{}
}
func (s *masterService) FilterStatusByID(id string, status []dto.Status) dto.Status {
	for _, data := range status {
		if data.ID == id {
			return data
		}
	}

	return dto.Status{}
}

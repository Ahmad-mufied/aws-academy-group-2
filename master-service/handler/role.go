package handler

import (
	"encoding/json"
	"errors"
	"master-service/model/entity"
	"master-service/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type RoleHandler interface {
	CreateNewRole(w http.ResponseWriter, r *http.Request)
	GetAllRoles(w http.ResponseWriter, r *http.Request)
	GetRoleByID(w http.ResponseWriter, r *http.Request)
}

type roleHandler struct {
	repo repository.RoleRepository
}

func (h *roleHandler) CreateNewRole(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}
	var role entity.Role

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Convert map back to struct
	requestBody, _ := json.Marshal(requestData)
	json.Unmarshal(requestBody, &role)

	// Ensure is_active defaults to true only if not provided in the request
	if _, exists := requestData["is_active"]; !exists {
		role.IsActive = true
	}

	// input validation
	if role.ID == uuid.Nil && role.Name == "" {
		http.Error(w, "Role name cannot be empty", http.StatusBadRequest)
		return
	}

	// save role to database
	if err := h.repo.CreateRole(&role); err != nil {
		http.Error(w, "Failed to create role", http.StatusInternalServerError)
		return
	}

	// get data by id
	newestRole, err := h.repo.FindRoleByID(role.ID)
	if err != nil {
		http.Error(w, "Failed to fetch updated status", http.StatusInternalServerError)
		return
	}

	// response success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newestRole); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *roleHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.repo.FindAllRoles()
	if err != nil {
		http.Error(w, "Failed to fetch roles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)

}

func (h *roleHandler) GetRoleByID(w http.ResponseWriter, r *http.Request) {
	// get id from params
	vars := mux.Vars(r)
	idParam := vars["id"]

	// convert string to uuid
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	// get role by id from database
	role, err := h.repo.FindRoleByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Role not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// response success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(role); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func InitRoleHandler(repo repository.RoleRepository) RoleHandler {
	return &roleHandler{repo: repo}
}

package handler

import (
	"encoding/json"
	"master-service/model/entity"
	"master-service/repository"
	"net/http"
)

type RoleHandler interface {
	CreateNewRole(w http.ResponseWriter, r *http.Request)
	GetAllRoles(w http.ResponseWriter, r *http.Request)
}

type roleHandler struct {
	repo repository.RoleRepository
}

func (h *roleHandler) CreateNewRole(w http.ResponseWriter, r *http.Request) {
	var role entity.Role

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateRole(&role); err != nil {
		http.Error(w, "Failed to create role", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}

func (h *roleHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles := h.repo.FindAllRoles()
	json.NewEncoder(w).Encode(roles)

}

func InitRoleHandler(repo repository.RoleRepository) RoleHandler {
	return &roleHandler{repo: repo}
}

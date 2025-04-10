package handler

import (
	"encoding/json"
	"errors"
	"master-service/dto"
	"master-service/model/entity"
	"master-service/repository"
	"net/http"
	"strings"
	"time"

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
	var req dto.CreateOrUpdateRoleRequest
	var role entity.Role

	// decode request ke struct DTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Cek validitas input
	if (req.ID == nil || *req.ID == uuid.Nil) && strings.TrimSpace(req.Name) == "" {
		http.Error(w, "Role name cannot be empty", http.StatusBadRequest)
		return
	}

	// Mapping ke entity.Role
	if req.ID != nil {
		role.ID = *req.ID
	}
	role.Name = req.Name

	if req.IsActive != nil {
		role.IsActive = *req.IsActive
	} else {
		role.IsActive = true
	}

	now := time.Now()

	if role.ID == uuid.Nil {
		// Create case
		role.ID = uuid.New()
		role.CreatedAt = now
		role.UpdatedAt = now

		if req.CreatedBy != nil {
			role.CreatedBy = *req.CreatedBy
			role.UpdatedBy = *req.CreatedBy
		}
	} else {
		// Update case
		role.UpdatedAt = now

		if req.UpdatedBy != nil {
			role.UpdatedBy = *req.UpdatedBy
		}
	}

	// Save via repo
	var err error
	if role.ID == uuid.Nil {
		err = h.repo.CreateRole(&role)
	} else {
		err = h.repo.UpdateRole(&role)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ambil data final dari DB
	newestRole, err := h.repo.FindRoleByID(role.ID)
	if err != nil {
		http.Error(w, "Failed to fetch updated role", http.StatusInternalServerError)
		return
	}

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

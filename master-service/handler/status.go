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

type StatusHandler interface {
	CreateNewStatus(w http.ResponseWriter, r *http.Request)
	GetAllStatus(w http.ResponseWriter, r *http.Request)
	GetStatusByID(w http.ResponseWriter, r *http.Request)
}

type statusHandler struct {
	repo repository.StatusRepository
}

func (h *statusHandler) CreateNewStatus(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateOrUpdateStatusRequest
	var status entity.Status

	// decode request ke struct DTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Cek validitas input
	if (req.ID == nil || *req.ID == uuid.Nil) && strings.TrimSpace(req.Name) == "" {
		http.Error(w, "Status name cannot be empty", http.StatusBadRequest)
		return
	}

	// Mapping ke entity.Status
	if req.ID != nil {
		status.ID = *req.ID
	}
	status.Name = req.Name

	if req.IsActive != nil {
		status.IsActive = *req.IsActive
	} else {
		status.IsActive = true
	}

	now := time.Now()

	if status.ID == uuid.Nil {
		// Create case
		status.ID = uuid.New()
		status.CreatedAt = now
		status.UpdatedAt = now

		if req.CreatedBy != nil {
			status.CreatedBy = *req.CreatedBy
			status.UpdatedBy = *req.CreatedBy
		}
	} else {
		// Update case
		status.UpdatedAt = now

		if req.UpdatedBy != nil {
			status.UpdatedBy = *req.UpdatedBy
		}
	}

	// Save via repo
	var err error
	if status.ID == uuid.Nil {
		err = h.repo.CreateStatus(&status)
	} else {
		err = h.repo.UpdateStatus(&status)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ambil data final dari DB
	newestStatus, err := h.repo.FindStatusByID(status.ID)
	if err != nil {
		http.Error(w, "Failed to fetch updated status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newestStatus); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *statusHandler) GetAllStatus(w http.ResponseWriter, r *http.Request) {
	status, err := h.repo.FindAllStatus()
	if err != nil {
		http.Error(w, "Failed to fetch status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func (h *statusHandler) GetStatusByID(w http.ResponseWriter, r *http.Request) {
	// get id from url params
	vars := mux.Vars(r)
	idParam := vars["id"]

	// convert string to uuid
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	// get status by id from database
	status, err := h.repo.FindStatusByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Status not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// response success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func InitStatusHandler(repo repository.StatusRepository) StatusHandler {
	return &statusHandler{repo: repo}
}

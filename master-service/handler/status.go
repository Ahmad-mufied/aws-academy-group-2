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

type StatusHandler interface {
	CreateNewStatus(w http.ResponseWriter, r *http.Request)
	GetAllStatus(w http.ResponseWriter, r *http.Request)
	GetStatusByID(w http.ResponseWriter, r *http.Request)
}

type statusHandler struct {
	repo repository.StatusRepository
}

func (h *statusHandler) CreateNewStatus(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}
	var status entity.Status

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Convert map back to struct
	requestBody, _ := json.Marshal(requestData)
	json.Unmarshal(requestBody, &status)

	// Ensure is_active defaults to true only if not provided in the request
	if _, exists := requestData["is_active"]; !exists {
		status.IsActive = true
	}

	// Input validation
	if status.ID == uuid.Nil && status.Name == "" {
		http.Error(w, "Status name cannot be empty", http.StatusBadRequest)
		return
	}

	// save status to database
	if err := h.repo.CreateStatus(&status); err != nil {
		http.Error(w, "Failed to create status", http.StatusInternalServerError)
		return
	}

	newestStatus, err := h.repo.FindStatusByID(status.ID)
	if err != nil {
		http.Error(w, "Failed to fetch updated status", http.StatusInternalServerError)
		return
	}

	// response success
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

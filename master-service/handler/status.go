package handler

import (
	"encoding/json"
	"master-service/model/entity"
	"master-service/repository"
	"net/http"
)

type StatusHandler interface {
	CreateNewStatus(w http.ResponseWriter, r *http.Request)
	GetAllStatus(w http.ResponseWriter, r *http.Request)
}

type statusHandler struct {
	repo repository.StatusRepository
}

func (h *statusHandler) CreateNewStatus(w http.ResponseWriter, r *http.Request) {
	var status entity.Status

	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateStatus(&status); err != nil {
		http.Error(w, "Failed to create status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(status)
}

func (h *statusHandler) GetAllStatus(w http.ResponseWriter, r *http.Request) {
	status := h.repo.FindAllStatus()
	json.NewEncoder(w).Encode(status)
}

func InitStatusHandler(repo repository.StatusRepository) StatusHandler {
	return &statusHandler{repo: repo}
}

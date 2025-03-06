package repository

import (
	"errors"
	"master-service/model/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusRepository interface {
	CreateStatus(status *entity.Status) error
	FindAllStatus() []entity.Status
}

type statusRepository struct {
	db *gorm.DB
}

func (r *statusRepository) CreateStatus(status *entity.Status) error {
	if status.Name == "" {
		return errors.New("status name cannot be empty")
	}

	if status.ID == uuid.Nil {
		status.ID = uuid.New()
		status.CreatedAt = time.Now()
	}

	return r.db.Omit("created_at").Save(&status).Error
}

func (r *statusRepository) FindAllStatus() []entity.Status {
	var status []entity.Status
	r.db.Find(&status)
	return status
}

func InitStatusRepository(db *gorm.DB) StatusRepository {
	return &statusRepository{db: db}
}

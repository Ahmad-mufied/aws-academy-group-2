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
	FindAllStatus() ([]entity.Status, error)
	FindStatusByID(id uuid.UUID) (*entity.Status, error)
}

type statusRepository struct {
	db *gorm.DB
}

func (r *statusRepository) CreateStatus(status *entity.Status) error {
	var existing entity.Status

	// Create new status
	if status.ID == uuid.Nil {
		if status.Name == "" {
			return errors.New("status name cannot be empty")
		}

		// Check if name already exists
		if err := r.db.Where("name = ?", status.Name).First(&existing).Error; err == nil {
			return errors.New("status name already exists")
		}

		// Assign new ID and timestamps
		status.ID = uuid.New()
		status.CreatedAt = time.Now()
		status.UpdatedAt = time.Now()

		return r.db.Create(&status).Error
	}

	// Update status
	if err := r.db.Where("id = ?", status.ID).First(&existing).Error; err != nil {
		return errors.New("status not found")
	}

	// Prevent name duplication with another existing status
	if status.Name != "" && status.Name != existing.Name {
		var duplicateCheck entity.Status
		if err := r.db.Where("name = ?", status.Name).First(&duplicateCheck).Error; err == nil {
			return errors.New("status name already exists")
		}
		existing.Name = status.Name
	}

	existing.IsActive = status.IsActive
	existing.UpdatedAt = time.Now()

	return r.db.Omit("created_at").Save(&existing).Error
}

func (r *statusRepository) FindAllStatus() ([]entity.Status, error) {
	var status []entity.Status
	result := r.db.Find(&status)
	if result.Error != nil {
		return nil, result.Error
	}
	return status, nil
}

func (r *statusRepository) FindStatusByID(id uuid.UUID) (*entity.Status, error) {
	var status entity.Status

	result := r.db.Take(&status, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &status, nil
}

func InitStatusRepository(db *gorm.DB) StatusRepository {
	return &statusRepository{db: db}
}

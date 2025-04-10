package repository

import (
	"errors"
	"master-service/model/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusRepository interface {
	CreateStatus(status *entity.Status) error
	UpdateStatus(status *entity.Status) error
	FindAllStatus() ([]entity.Status, error)
	FindStatusByID(id uuid.UUID) (*entity.Status, error)
}

type statusRepository struct {
	db *gorm.DB
}

func (r *statusRepository) CreateStatus(status *entity.Status) error {
	return r.db.Create(status).Error
}

func (r *statusRepository) UpdateStatus(status *entity.Status) error {
	var existing entity.Status
	if err := r.db.First(&existing, "id = ?", status.ID).Error; err != nil {
		return err
	}

	// Cek duplikasi name jika nama berubah
	if status.Name != existing.Name {
		var duplicate entity.Status
		if err := r.db.First(&duplicate, "name = ?", status.Name).Error; err == nil {
			return errors.New("status name already exists")
		}
	}

	existing.Name = status.Name
	existing.IsActive = status.IsActive
	existing.UpdatedBy = status.UpdatedBy

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

package repository

import (
	"errors"
	"master-service/model/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *entity.Role) error
	FindAllRoles() ([]entity.Role, error)
	FindRoleByID(id uuid.UUID) (*entity.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) CreateRole(role *entity.Role) error {
	var existing entity.Role

	// Create new role
	if role.ID == uuid.Nil {
		if role.Name == "" {
			return errors.New("role name cannot be empty")
		}

		// Check if name already exists
		if err := r.db.Where("name = ?", role.Name).First(&existing).Error; err == nil {
			return errors.New("role name already exists")
		}

		// Assign new ID and timestamps
		role.ID = uuid.New()
		role.CreatedAt = time.Now()
		role.UpdatedAt = time.Now()

		return r.db.Create(&role).Error
	}

	// Update role
	if err := r.db.Where("id = ?", role.ID).First(&existing).Error; err != nil {
		return errors.New("role not found")
	}

	// Prevent name duplication with another existing role
	if role.Name != "" && role.Name != existing.Name {
		var duplicateCheck entity.Status
		if err := r.db.Where("name = ?", role.Name).First(&duplicateCheck).Error; err == nil {
			return errors.New("role name already exists")
		}
		existing.Name = role.Name
	}

	existing.IsActive = role.IsActive
	existing.UpdatedAt = time.Now()

	return r.db.Omit("created_at").Save(&existing).Error
}

func (r *roleRepository) FindAllRoles() ([]entity.Role, error) {
	var roles []entity.Role
	result := r.db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (r *roleRepository) FindRoleByID(id uuid.UUID) (*entity.Role, error) {
	var role entity.Role

	result := r.db.Take(&role, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &role, nil
}

func InitRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

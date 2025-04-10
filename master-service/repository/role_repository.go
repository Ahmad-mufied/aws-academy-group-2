package repository

import (
	"errors"
	"master-service/model/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *entity.Role) error
	UpdateRole(role *entity.Role) error
	FindAllRoles() ([]entity.Role, error)
	FindRoleByID(id uuid.UUID) (*entity.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) CreateRole(role *entity.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) UpdateRole(role *entity.Role) error {
	var existing entity.Role
	if err := r.db.First(&existing, "id = ?", role.ID).Error; err != nil {
		return err
	}

	// Cek duplikasi name jika nama berubah
	if role.Name != existing.Name {
		var duplicate entity.Role
		if err := r.db.First(&duplicate, "name = ?", role.Name).Error; err == nil {
			return errors.New("role name already exists")
		}
	}

	existing.Name = role.Name
	existing.IsActive = role.IsActive
	existing.UpdatedBy = role.UpdatedBy

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

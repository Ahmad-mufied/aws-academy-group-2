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
	FindAllRoles() []entity.Role
}

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) CreateRole(role *entity.Role) error {
	// Validation
	if role.Name == "" {
		return errors.New("role name cannot be empty")
	}

	if role.ID == uuid.Nil {
		role.ID = uuid.New()
		role.CreatedAt = time.Now()
	}

	return r.db.Omit("created_at").Save(&role).Error
}

func (r *roleRepository) FindAllRoles() []entity.Role {
	var roles []entity.Role
	r.db.Find(&roles)
	return roles
}

func InitRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

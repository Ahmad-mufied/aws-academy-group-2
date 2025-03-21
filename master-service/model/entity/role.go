package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;column:id"`
	Name      string    `json:"name" gorm:"column:name;not null;unique"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active"`
	CreatedBy uuid.UUID `json:"created_by" gorm:"type:uuid;column:created_by"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedBy uuid.UUID `json:"updated_by" gorm:"type:uuid;column:updated_by"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (Role) TableName() string {
	return "roles"
}

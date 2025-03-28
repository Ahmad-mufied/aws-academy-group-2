package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `gorm:"type:char(36);primary_key"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Email       string         `gorm:"type:varchar(255);unique;not null"`
	ProductIDs  datatypes.JSON `gorm:"type:json;not null"`
	RoleID      uuid.UUID      `gorm:"type:char(36);not null"`
	StatusID    uuid.UUID      `gorm:"type:char(36);not null"`
	DateOfBirth time.Time      `gorm:"column:date_of_birth;type:date;not null"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
}

func (User) TableName() string { return "users" }

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

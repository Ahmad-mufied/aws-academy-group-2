package repository

import (
	"errors"

	"github.com/DavidAfdal/user-services/internal/model"
	"gorm.io/gorm"
)

var ErrExitedUser = errors.New("email already taken")

type UserRepository interface {
	GetUsers(search string, page, limit int) ([]model.User, int, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers(search string, page, limit int) ([]model.User, int, error) {
	var total int64
	users := make([]model.User, 0)

	query := r.db.Model(&model.User{})

	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Offset((page - 1) * limit).Limit(limit).Find(&users); r.db.Error != nil {
		return nil, 0, r.db.Error
	}

	return users, int(total), nil
}

func (r *userRepository) CreateUser(user *model.User) (*model.User, error) {
	err := r.db.Create(user).Error

	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return nil, ErrExitedUser
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(user *model.User) error {
	return r.db.Delete(user).Error
}

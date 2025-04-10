package repository

import (
	"encoding/json"
	"strings"

	"github.com/DavidAfdal/user-services/internal/model"
	"github.com/DavidAfdal/user-services/pkg/constant"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(search string, page, limit int, startDate, endDate string) ([]model.User, int, error)
	GetUser(id string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id string) error
	AssignProducts(user *model.User, productIDs []string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers(search string, page, limit int, startDate, endDate string) ([]model.User, int, error) {
	var total int64
	users := make([]model.User, 0)

	query := r.db.Model(&model.User{})

	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if startDate != "" && endDate != "" {
		query = query.Where("date_of_birth BETWEEN ? AND ?", startDate, endDate)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.Offset((page - 1) * limit).Limit(limit).Find(&users); r.db.Error != nil {
		return nil, 0, r.db.Error
	}

	return users, int(total), nil
}

func (r *userRepository) GetUser(id string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) CreateUser(user *model.User) (*model.User, error) {
	err := r.db.Create(user).Error

	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, constant.ErrUserExists
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

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *userRepository) AssignProducts(user *model.User, productIDs []string) error {
	jsonData, err := json.Marshal(productIDs)

	if err != nil {
		return err
	}

	user.ProductIDs = jsonData

	return r.db.Save(user).Error
}

package repository

import (
	"github.com/tersorasora/freelance/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}

type userRepositoryData struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryData{db}
}

func (urd *userRepositoryData) CreateUser(user *entity.User) error {
	return urd.db.Create(user).Error
}

func (urd *userRepositoryData) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := urd.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
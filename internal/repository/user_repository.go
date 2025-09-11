package repository

import (
	"github.com/tersorasora/freelance/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	DeleteUser(userID string) error
	GetUserByEmail(email string) (*entity.User, error)
	GetLastUserID() (string, error)
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

func (urd *userRepositoryData) DeleteUser(userID string) error {
	return urd.db.Where("user_id = ?", userID).Delete(&entity.User{}).Error
}

func (urd *userRepositoryData) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := urd.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (urd *userRepositoryData) GetLastUserID() (string, error) {
	var lastUser string

	err := urd.db.Model(&entity.User{}).Order("user_id DESC").Limit(1).Pluck("user_id", &lastUser).Error
	if err != nil{
		return "", err
	}

	return lastUser, nil
}
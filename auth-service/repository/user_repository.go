package repository

import (
	model "auth-service/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Register(user *model.User) (bool, error) {
	var exist model.User
	res := r.DB.Where("username = ?", user.Username).First(&exist).Error
	return res == nil, r.DB.Create(user).Error
}

func (r *UserRepository) Login(username string) (model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

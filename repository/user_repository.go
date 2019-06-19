package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/youtangai/cloud-training/model"
)

type IUserRepository interface {
	CreateUser(user model.User) error
	FindUserByID(userID string) (*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) CreateUser(user model.User) error {
	db := repo.db

	err := db.Create(&user).Error

	if err != nil {
		return fmt.Errorf("repository:create user error:%s", err)
	}

	return nil
}

func (repo UserRepository) FindUserByID(userID string) (*model.User, error) {
	db := repo.db

	var user model.User
	err := db.First(&user, "user_id = ?", userID).Error

	if err != nil {
		return nil, fmt.Errorf("repository:find user error. user_id = %s, err = %s", userID, err)
	}

	return &user, nil
}

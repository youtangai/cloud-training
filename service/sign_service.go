package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/youtangai/cloud-training/model"
	"github.com/youtangai/cloud-training/repository"
)

type ISignService interface {
	GetAccessToken(user model.User) (string, error)
	SignUpUser(user model.User) error
}

type SignService struct {
	repo repository.IUserRepository
}

func NewSignService(repo repository.IUserRepository) ISignService {
	return SignService{
		repo: repo,
	}
}

func (srv SignService) GetAccessToken(user model.User) (string, error) {
	result, err := srv.repo.FindUserByID(user.UserID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return "", fmt.Errorf("service:err: no user found in db. userid=%s, err=%s", user.UserID, err)
		default:
			return "", fmt.Errorf("service:err: unknown error occured. err=%s",err)
		}
	}

	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))
	if result.Password != hash {
		return "", fmt.Errorf("service:err: not match password. base=%s, input=%s", result.Password, hash)
	}

	return result.AccessToken, nil
}

func (srv SignService) SignUpUser(user model.User) error {
	//password hashing
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))

	//update password hashing password
	user.Password = hash

	//generate access token
	token := fmt.Sprintf("%x", sha256.Sum256([]byte(user.UserID+user.Password)))
	user.AccessToken = token

	//create user
	err := srv.repo.CreateUser(user)

	if err != nil {
		return fmt.Errorf("service:err: failed to create user. user=%v, err=%s", user, err)
	}

	return nil
}

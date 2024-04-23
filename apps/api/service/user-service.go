package service

import (
	"github.com/timzolleis/smallstatus/helper"
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service *UserService) CreateUser(name string, email string, password string) (*model.User, error) {

	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := model.User{
		Email:    email,
		Name:     name,
		Password: hashedPassword,
	}
	createdUser, err := service.Repository.Create(&user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (service *UserService) FindUserById(id int) (*model.User, error) {
	user, err := service.Repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) FindAll() ([]model.User, error) {
	users, err := service.Repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *UserService) Delete(id int) error {
	err := service.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

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
	return service.Repository.Create(&user)
}

func (service *UserService) FindUserById(id uint) (*model.User, error) {
	return service.Repository.FindById(id)
}

func (service *UserService) FindUserByEmail(email string) (*model.User, error) {
	return service.Repository.FindByEmail(email)
}

func (service *UserService) FindAll() ([]model.User, error) {
	return service.Repository.FindAll()
}

func (service *UserService) Delete(id int) error {
	return service.Repository.Delete(id)
}

package service

import (
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service *UserService) Create(user *model.User) (*model.User, error) {
	return service.Repository.Create(user)
}

func (service *UserService) FindById(id uint) (*model.User, error) {
	return service.Repository.FindById(id)
}

func (service *UserService) FindByEmail(email string) (*model.User, error) {
	return service.Repository.FindByEmail(email)
}

func (service *UserService) FindAll() ([]model.User, error) {
	return service.Repository.FindAll()
}

func (service *UserService) Delete(id int) error {
	return service.Repository.Delete(id)
}

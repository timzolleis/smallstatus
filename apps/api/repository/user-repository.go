package repository

import (
	"github.com/timzolleis/smallstatus/database"
	"github.com/timzolleis/smallstatus/model"
)

type UserRepository struct{}

func (repo *UserRepository) Create(user *model.User) (*model.User, error) {
	err := database.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) FindById(id int) (*model.User, error) {
	var user model.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Delete(id int) error {
	err := database.DB.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

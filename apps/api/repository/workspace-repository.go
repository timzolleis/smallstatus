package repository

import (
	"github.com/timzolleis/smallstatus/database"
	"github.com/timzolleis/smallstatus/model"
	"gorm.io/gorm"
)

type WorkspaceRepository struct {
}

func (repo *WorkspaceRepository) FindById(id uint) (*model.Workspace, error) {
	var workspace model.Workspace
	err := database.DB.Where("id = ?", id).Find(&workspace).Error
	return &workspace, err
}

func (repo *WorkspaceRepository) FindAll() ([]model.Workspace, error) {
	var workspaces []model.Workspace
	err := database.DB.Find(&workspaces).Error
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (repo *WorkspaceRepository) Create(workspace *model.Workspace) (*model.Workspace, error) {
	err := database.DB.Create(workspace).Error
	return workspace, err
}

func (repo *WorkspaceRepository) Update(workspace *model.Workspace) (*model.Workspace, error) {
	result := database.DB.Save(workspace)
	err := result.Error
	if result.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return workspace, err
}

func (repo *WorkspaceRepository) Delete(id uint) error {
	result := database.DB.Delete(&model.Workspace{}, id)
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (repo *WorkspaceRepository) FindByUserId(userId uint) ([]model.Workspace, error) {
	var workspaces []model.Workspace
	userRepo := UserRepository{}
	user, err := userRepo.FindById(userId)
	if err != nil || user == nil {
		return nil, err
	}
	err = database.DB.Model(&user).Association("Workspaces").Find(&workspaces)
	return workspaces, err
}

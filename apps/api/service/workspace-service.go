package service

import (
	"github.com/timzolleis/smallstatus/model"
	"github.com/timzolleis/smallstatus/repository"
)

type WorkspaceService struct {
	repository repository.WorkspaceRepository
}

func (s *WorkspaceService) FindById(id uint) (*model.Workspace, error) {
	return s.repository.FindById(id)
}

func (s *WorkspaceService) FindUserWorkspaces(userId uint) ([]model.Workspace, error) {
	return s.repository.FindByUserId(userId)
}

func (s *WorkspaceService) IsPartOfWorkspace(userId uint, workspaceId uint) bool {
	userWorkspaces, err := s.repository.FindByUserId(userId)
	if err != nil {
		return false
	}
	for _, workspace := range userWorkspaces {
		if workspace.ID == workspaceId {
			return true
		}
	}
	return false
}

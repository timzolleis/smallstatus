package service

import "github.com/timzolleis/smallstatus/repository"

type WorkspaceService struct {
	repository repository.WorkspaceRepository
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

package model

type User struct {
	Base
	Name       string      `json:"name"`
	Email      string      `json:"email" gorm:"uniqueIndex"`
	Password   string      `json:"password"`
	Workspaces []Workspace `gorm:"many2many:workspace_users;" json:"workspaces"`
}

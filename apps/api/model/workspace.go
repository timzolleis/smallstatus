package model

type Workspace struct {
	Base
	Name     string    `json:"name"`
	ApiKeys  []ApiKey  `json:"api_keys"`
	Users    []User    `gorm:"many2many:workspace_users;" json:"users"`
	Monitors []Monitor `json:"monitors"`
}

package model

import "gorm.io/gorm"

type Monitor struct {
	gorm.Model
	Name        string
	Url         string
	Type        string
	Interval    int
	WorkspaceID uint
	Headers     []MonitorHeader `gorm:"foreignKey:MonitorID"`
}

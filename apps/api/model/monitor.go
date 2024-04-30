package model

import "gorm.io/gorm"

type Monitor struct {
	gorm.Model
	Name        string
	Url         string
	Interval    int
	Timeout     int
	Retries     int
	Method      string
	WorkspaceID uint
	Headers     []MonitorHeader `gorm:"foreignKey:MonitorID"`
}

package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint           `gorm:"primarykey,<-:create" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

package models

import "gorm.io/gorm"

type Monitor struct {
	gorm.Model
	Name     string `json:"name"`
	Url      string `json:"url"`
	Interval int    `json:"interval"`
}

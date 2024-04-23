package models

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	Name    string
	Value   string
	revoked bool
}

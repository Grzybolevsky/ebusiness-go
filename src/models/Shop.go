package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name string `json:"Name,omitempty" gorm:"not null;"`
}

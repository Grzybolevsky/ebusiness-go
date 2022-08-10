package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string `json:"FirstName"`
	SecondName string `json:"SecondName"`
	Email      string `json:"Email" gorm:"unique; not null;"`
}

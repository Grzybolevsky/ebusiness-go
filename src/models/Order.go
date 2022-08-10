package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Code     string    `json:"Code,omitempty"`
	UserId   int       `json:"UserId,omitempty" gorm:"index; not null;"`
	User     User      `json:"User"`
	ShopId   int       `json:"ShopId,omitempty" gorm:"index; not null;"`
	Shop     Shop      `json:"Shop"`
	Date     time.Time `json:"Date"`
	Finished bool      `json:"Finished"`
}

package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code     string  `json:"Code"`
	Name     string  `json:"Name"`
	Price    float64 `json:"Price"`
	Quantity uint32  `json:"Quantity" gorm:"check: Quantity > 0"`
	ShopId   int     `json:"ShopId" gorm:"index; not null;"`
	Shop     Shop    `json:"Shop"`
}

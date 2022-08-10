package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Quantity   int     `json:"Quantity,omitempty"`
	ProductId  int     `json:"ProductId,omitempty" gorm:"index; not null;"`
	Product    Product `json:"Product"`
	OrderId    int     `json:"OrderId,omitempty" gorm:"index; not null;"`
	Order      Order   `json:"Order"`
	TotalPrice float64 `json:"TotalPrice,omitempty"`
}

package models

import "time"

type Order struct{
	GormModel
	OrderedAt 		*time.Time 	`gorm:"not null;" json:"order_at" form:"order_at" valid:"required"`
	CustomerName 	string 		`gorm:"not null;" json:"customer_name" form:"customer_name" valid:"required"`
	Items 			[]Items 	`gorm:"foreignKey:OrderID"`
}

type Items struct{
	GormModel
	OrderID			uint	`gorm:"not null;" json:"order_id"`
	ItemCode 		string	`gorm:"not null;" json:"item_code"`
	Description 	string	`gorm:"not null;" json:"description"`
	Quantity		int	`gorm:"not null;" json:"quantity"`
}
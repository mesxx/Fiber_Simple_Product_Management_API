package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	UserId      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}

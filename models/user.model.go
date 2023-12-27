package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `gorm:"unique" json:"name"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"password"`
	Products []Product `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

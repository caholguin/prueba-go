package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"not null" validate:"required,min=3"`
	Email string `json:"email" gorm:"not null;unique" validate:"required,email"`
}
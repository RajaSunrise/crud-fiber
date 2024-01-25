package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama string `json:"nama"`
	Email string `json:"email"`
	Umur int8 `json:"umur"`
}
package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

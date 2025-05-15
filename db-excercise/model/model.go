package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey:autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null;unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Food      []Food    `gorm:"foreignKey:CategoryID;references:ID"`
}

type Food struct {
	gorm.Model
	ID           uint         `json:"id" gorm:"primaryKey:autoIncrement"`
	CategoryID   uint         `json:"category_id" gorm:"not null"`
	IngredientID uint         `json:"ingredient_id" gorm:"not null"`
	Name         string       `json:"name" gorm:"type:varchar(100);not null"`
	Price        float64      `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Ingredient   []Ingredient `gorm:"many2many:food_ingredients"`
}

type Ingredient struct {
	gorm.Model
	ID         uint       `json:"id" gorm:"primaryKey:autoIncrement"`
	CategoryID uint       `json:"category_id" gorm:"not null"`
	FoodID     uint       `json:"food_id" gorm:"not null"`
	Name       string     `json:"name" gorm:"type:varchar(100);not null"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Category   []Category `gorm:"many2many:category_ingredients;foreignKey:CategoryID"`
	Food       []Food     `gorm:"many2many:food_ingredients"`
}

package model

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name           string           `json:"name"`
	Price          uint             `json:"price"`
	FoodIngredient []FoodIngredient `gorm:"foreignKey:FoodID"`
	CategoryID     uint             `json:"category_id"` // Must match Category.ID (uint)
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Food []Food `gorm:"foreignKey:CategoryID"`
}

type FoodIngredient struct {
	gorm.Model
	FoodID         uint   `json:"food_id"` // Must match Food.ID (uint)
	FoodName       string `json:"food_name"`
	IngredientID   uint   `json:"ingredient_id"` // Must match Ingredient.ID (uint)
	IngredientName string `json:"ingredient_name"`
}

type Ingredient struct {
	gorm.Model
	Name           string           `json:"name"`
	FoodIngredient []FoodIngredient `gorm:"foreignKey:IngredientID"`
}

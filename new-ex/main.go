package main

import (
	"db-excercise/conf"
	"db-excercise/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstant *gorm.DB

type FoodIngredientCount struct {
	CategoryID   uint   `json:"cat_id"`
	CategoryName string `json:"cat_name"`
	FoodCount    int64  `json:"food_count"`
}

func getListFoodIngredients(c *gin.Context) {
	var foodIngredients []model.FoodIngredient
	if err := dbInstant.Model(&model.FoodIngredient{}).Find(&foodIngredients).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch food ingredients"})
		return
	}
	c.JSON(200, foodIngredients)
}

func getCountOfEachCategory(c *gin.Context) {
	var cats []model.Category
	if err := dbInstant.Model(&model.Category{}).Find(&cats).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch categories"})
		return
	}
	var res []FoodIngredientCount = []FoodIngredientCount{}

	for _, cat := range cats {
		var count int64
		if err := dbInstant.Model(&model.Food{}).Where("category_id = ?", cat.ID).Count(&count).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to count food"})
			return
		}
		res = append(res, FoodIngredientCount{
			CategoryID:   cat.ID,
			CategoryName: cat.Name,
			FoodCount:    count,
		})
	}
	c.JSON(200, res)
}

func getCountOfEachCategoryMoreComplexQuery(c *gin.Context) {
	var res []FoodIngredientCount
	if err := dbInstant.Table("categories").Select("categories.id as cat_id, categories.name as cat_name, count(foods.id) as food_count").
		Joins("left join foods on categories.id = foods.category_id").
		Group("categories.id").
		Scan(&res).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch food count"})
		return
	}
	c.JSON(200, res)
}

type FoodNoNeed struct {
	FoodID   uint   `json:"food_id"`
	FoodName string `json:"food_name"`
}

type NoNeedRes struct {
	CategoryID   uint         `json:"cat_id"`
	CategoryName string       `json:"cat_name"`
	Food         []FoodNoNeed `json:"foods"`
}

func NoNeed(c *gin.Context) {
	var res []model.Category
	if err := dbInstant.Model(&model.Category{}).Where("name NOT IN ?", []string{"egg", "miso"}).Find(&res).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch categories"})
		return
	}
	var foodNoNeed []FoodNoNeed

	var food1 []model.Food
	var NoNeedResvar []NoNeedRes
	for _, cat := range res {
		if err := dbInstant.Model(&model.Food{}).Where("category_id = ?", cat.ID).Find(&food1).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch food items"})
			return
		}
		for _, food := range food1 {
			foodNoNeed = append(foodNoNeed, FoodNoNeed{
				FoodID:   food.ID,
				FoodName: food.Name,
			})
		}
		NoNeedResvar = append(NoNeedResvar, NoNeedRes{
			CategoryID:   cat.ID,
			CategoryName: cat.Name,
			Food:         foodNoNeed,
		})
	}
	c.JSON(200, NoNeedResvar)
}

func NoNeedEgg(c *gin.Context) {
	var res []model.Category
	if err := dbInstant.Model(&model.Category{}).Where("name NOT IN ?", []string{"egg"}).Find(&res).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch categories"})
		return
	}
	var foodNoNeed []FoodNoNeed

	var food1 []model.Food
	var NoNeedResvar []NoNeedRes
	for _, cat := range res {
		if err := dbInstant.Model(&model.Food{}).Where("category_id = ?", cat.ID).Find(&food1).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch food items"})
			return
		}
		for _, food := range food1 {
			foodNoNeed = append(foodNoNeed, FoodNoNeed{
				FoodID:   food.ID,
				FoodName: food.Name,
			})
		}
		NoNeedResvar = append(NoNeedResvar, NoNeedRes{
			CategoryID:   cat.ID,
			CategoryName: cat.Name,
			Food:         foodNoNeed,
		})
	}
	c.JSON(200, NoNeedResvar)
}

func NoNeedMiso(c *gin.Context) {
	var res []model.Category
	if err := dbInstant.Model(&model.Category{}).Where("name NOT IN ?", []string{"miso"}).Find(&res).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch categories"})
		return
	}
	var foodNoNeed []FoodNoNeed

	var food1 []model.Food
	var NoNeedResvar []NoNeedRes
	for _, cat := range res {
		if err := dbInstant.Model(&model.Food{}).Where("category_id = ?", cat.ID).Find(&food1).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch food items"})
			return
		}
		for _, food := range food1 {
			foodNoNeed = append(foodNoNeed, FoodNoNeed{
				FoodID:   food.ID,
				FoodName: food.Name,
			})
		}
		NoNeedResvar = append(NoNeedResvar, NoNeedRes{
			CategoryID:   cat.ID,
			CategoryName: cat.Name,
			Food:         foodNoNeed,
		})
	}
	c.JSON(200, NoNeedResvar)
}

func createCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := dbInstant.Create(&category).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(201, category)
}

func main() {
	config, err := conf.NewConfig()
	if err != nil {
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	fmt.Printf("DSN = %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbInstant = db
	server := gin.Default()
	server.POST("/category", createCategory)
	server.GET("/food_ingredients", getListFoodIngredients)
	server.GET("/count_of_each_category", getCountOfEachCategoryMoreComplexQuery)
	server.GET("/no_need", NoNeed)
	server.GET("/no_need_egg", NoNeedEgg)
	server.GET("/no_need_miso", NoNeedMiso)
	server.Run(":8080")
}

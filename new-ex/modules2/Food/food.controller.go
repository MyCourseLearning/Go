package food

import "github.com/gin-gonic/gin"

type FoodController interface {
	GetAll(c *gin.Context)
}

type foodController struct {
	foodService FoodService
}

func NewFoodController(foodService FoodService) FoodController {
	return &foodController{
		foodService: foodService,
	}
}

func (f *foodController) GetAll(c *gin.Context) {
	foods, err := f.foodService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch foods"})
		return
	}
	c.JSON(200, gin.H{
		"foods": foods,
	})
}

package main

import (
	"fmt"
	"test/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SubTopic struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Amount int    `json:"amount"`
}

type UpdateSupTopic struct {
	Title  string `json:"title"`
	Amount int    `json:"amount"`
}

type Response struct {
	ID       string     `json:"id"`
	SubTopic []SubTopic `json:"sub_topic"`
}

var data []Response = []Response{
	{
		ID: "1",
		SubTopic: []SubTopic{
			{
				ID:     "1",
				Title:  "Sub Topic 1",
				Amount: 10,
			},
			{
				ID:     "2",
				Title:  "Sub Topic 2",
				Amount: 20,
			},
		},
	},
	{
		ID: "2",
		SubTopic: []SubTopic{
			{
				ID:     "3",
				Title:  "Sub Topic 3",
				Amount: 30,
			},
			{
				ID:     "4",
				Title:  "Sub Topic 4",
				Amount: 40,
			},
		},
	},
}

func getHello(c *gin.Context) {
	c.JSON(200, data)
}

func getByID(c *gin.Context) {
	id := c.Param("id")
	for _, r := range data {
		if r.ID == id {
			c.JSON(200, r)
			return
		}
	}
	c.JSON(404, gin.H{"message": "not found"})
}

func postHello(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"message": "id is required"})
		return
	}
	var newSubTopic SubTopic
	if err := c.ShouldBindJSON(&newSubTopic); err != nil {
		c.JSON(400, gin.H{"message": "invalid input"})
		return
	}
	for _, r := range data {
		if id == r.ID {
			r.SubTopic = append(r.SubTopic, newSubTopic)
			c.JSON(200, r)
			return
		}
	}
	c.JSON(404, gin.H{"message": "not found"})
}

func updateHello(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"message": "id is required"})
		return
	}
	subID := c.Param("sub_id")
	if subID == "" {
		c.JSON(400, gin.H{"message": "sub_id is required"})
		return
	}
	var newSubTopic UpdateSupTopic
	if err := c.ShouldBindJSON(&newSubTopic); err != nil {
		c.JSON(400, gin.H{"message": "invalid input"})
		return
	}
	var index int = -1
	for i, r := range data {
		if id == r.ID {
			index = i
		}
	}
	if index == -1 {
		c.JSON(404, gin.H{"message": "not found"})
		return
	}
	for i, sub := range data[index].SubTopic {
		if sub.ID == subID {
			sub.Title = newSubTopic.Title
			sub.Amount = newSubTopic.Amount
			data[index].SubTopic[i] = sub
			c.JSON(200, data[index])
			return
		}
	}
	c.JSON(404, gin.H{"message": "sub topic not found"})
}

func deleteHello(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"message": "id is required"})
		return
	}
	subID := c.Param("sub_id")
	if subID == "" {
		c.JSON(400, gin.H{"message": "sub_id is required"})
		return
	}
	var index int = -1
	for i, r := range data {
		if id == r.ID {
			index = i
		}
	}
	if index == -1 {
		c.JSON(404, gin.H{"message": "not found"})
		return
	}
	for i, sub := range data[index].SubTopic {
		if sub.ID == subID {
			data[index].SubTopic = append(data[index].SubTopic[:i], data[index].SubTopic[i+1:]...)
			c.JSON(200, data[index])
			return
		}
	}
	c.JSON(404, gin.H{"message": "sub topic not found"})
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var dbInstant *gorm.DB

func getUsers(c *gin.Context) {
	var users []model.UserModel
	if err := dbInstant.Model(&model.UserModel{}).Find(&users).Error; err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "failed to get users"})
		return
	}
	var userResponse []UserResponse = []UserResponse{}
	for _, user := range users {
		userResponse = append(userResponse, UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}
	c.JSON(200, userResponse)
}

func main() {
	dsn := "user:user_password@tcp(localhost:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbInstant = db
	// db.Migrator().DropTable(&model.UserModel{})
	// db.AutoMigrate(&model.UserModel{})

	server := gin.Default()
	// server.GET("/hello", getHello)
	// server.GET("/hello/:id", getByID)
	// server.POST("/hello/:id", postHello)
	// server.PUT("/hello/:id/:sub_id", updateHello)
	// server.DELETE("/hello/:id/:sub_id", deleteHello)
	server.GET("user", getUsers)
	server.Run(":8080")
}

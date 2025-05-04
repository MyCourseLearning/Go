package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SubTopic struct {
	ID     string `json:"id"`
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
	var body SubTopic
	test := c.ShouldBindBodyWithJSON(&body)
	fmt.Println("test", test)
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	fmt.Println("Received:", body)
	data[1].SubTopic = append(data[1].SubTopic, SubTopic{
		ID:     body.ID,
		Title:  body.Title,
		Amount: body.Amount,
	})
	c.JSON(200, data)
}

func main() {
	server := gin.Default()
	server.GET("/hello", getHello)
	server.GET("/hello/:id", getByID)
	server.POST("/hello", postHello)
	server.Run(":8080")
}

package main

import "github.com/gin-gonic/gin"

type SubTopic struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Amount int    `json:"amount"`
}

type Response struct {
	ID       string     `json:"id"`
	SubTopic []SubTopic `json:"sub_topic"`
}

var res []Response = []Response{
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
	c.JSON(200, res)
}

func getByID(c *gin.Context) {
	id := c.Param("id")
	for _, r := range res {
		if r.ID == id {
			c.JSON(200, r)
			return
		}
	}
	c.JSON(404, gin.H{"message": "not found"})
}

func main() {
	server := gin.Default()
	server.GET("/hello", getHello)
	server.GET("/hello/:id", getByID)
	server.Run(":8080")
}

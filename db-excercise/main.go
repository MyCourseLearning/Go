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

func main() {
	config, err := conf.NewConfig()
	if err != nil {
		return
	}
	// dsn := "user:user_password@tcp(localhost:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	fmt.Printf("DSN = %s", dsn)
	// return
	// return
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbInstant = db
	db.Migrator().DropTable(&model.Category{}, &model.Food{}, &model.Ingredient{})
	db.AutoMigrate(&model.Category{}, &model.Food{}, &model.Ingredient{})

	server := gin.Default()
	// server.GET("/hello", getHello)
	// server.GET("/hello/:id", getByID)
	// server.POST("/hello/:id", postHello)
	// server.PUT("/hello/:id/:sub_id", updateHello)
	// server.DELETE("/hello/:id/:sub_id", deleteHello)
	server.Run(":8080")
}

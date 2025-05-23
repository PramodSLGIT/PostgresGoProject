package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Books struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetUpRoutes(api *gin.RouterGroup) {
	api.Group("/api")
	api.POST("?addbooks")
	api.GET("/getbooks")
	api.GET("/getbook/:id")
	api.DELETE("/deletebook")
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	fmt.Println("Hello Postgres")
	router.Run(":3010")
}

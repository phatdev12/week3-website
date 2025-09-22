package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phatdev12/week3-web/database"
	"github.com/phatdev12/week3-web/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connect()

	r := gin.Default()
	r.LoadHTMLGlob("app/*")

	r.GET("/", func(c *gin.Context) {
		var categories []models.Category
		database.DB.Find(&categories)

		c.HTML(200, "index.tmpl", gin.H{
			"categories": categories,
		})
	})

	r.GET("/category/:id", func(c *gin.Context) {
		var products []models.Product
		database.DB.Where("id_category = ?", c.Param("id")).Find(&products)

		c.HTML(200, "category.tmpl", gin.H{
			"products": products,
		})
	})
	r.Run(":3000")
}

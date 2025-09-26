package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phatdev12/week3-website/database"
	"github.com/phatdev12/week3-website/models"
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

	r.GET("/category", func(c *gin.Context) {
		var products []models.Product
		database.DB.Find(&products)

		c.HTML(200, "category.tmpl", gin.H{
			"products": products,
		})
	})

	r.GET("/category/:id", func(c *gin.Context) {
		var products []models.Product
		database.DB.Where("id_category = ?", c.Param("id")).Find(&products)

		c.HTML(200, "category.tmpl", gin.H{
			"products": products,
		})
	})

	r.GET("/create-product", func(c *gin.Context) {
		var categories []models.Category
		database.DB.Find(&categories)

		c.HTML(200, "createProduct.tmpl", gin.H{
			"categories": categories,
		})
	})

	r.POST("/products", func(c *gin.Context) {
		var product models.Product

		if err := c.ShouldBind(&product); err != nil {
			c.String(400, "Bad request")
			return
		}

		database.DB.Create(&product)
		c.Redirect(302, "/")
	})

	r.Run(":3000")
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"strconv"

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

	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"formatPrice": func(price float64) string {
			return fmt.Sprintf("%.0f VNĐ", price)
		},
		"formatPriceWithCommas": func(price float64) string {
			priceInt := int(price)
			priceStr := fmt.Sprintf("%d", priceInt)

			n := len(priceStr)
			if n <= 3 {
				return priceStr + " VNĐ"
			}

			var result []byte
			for i, digit := range []byte(priceStr) {
				if i > 0 && (n-i)%3 == 0 {
					result = append(result, ',')
				}
				result = append(result, digit)
			}

			return string(result) + " VNĐ"
		},
		"pageRange": func(current, total int) []int {
			start := current - 2
			if start < 1 {
				start = 1
			}
			end := current + 2
			if end > total {
				end = total
			}

			var pages []int
			for i := start; i <= end; i++ {
				pages = append(pages, i)
			}
			return pages
		},
	})

	r.LoadHTMLGlob("app/*")
	r.Static("/static", "./static")

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
		var totalCount int64

		page := 1
		limit := 8

		if p := c.Query("page"); p != "" {
			if parsedPage, err := strconv.Atoi(p); err == nil && parsedPage > 0 {
				page = parsedPage
			}
		}

		if l := c.Query("limit"); l != "" {
			if parsedLimit, err := strconv.Atoi(l); err == nil && parsedLimit > 0 && parsedLimit <= 100 {
				limit = parsedLimit
			}
		}

		offset := (page - 1) * limit
		categoryID := c.Param("id")

		database.DB.Model(&models.Product{}).Where("id_category = ?", categoryID).Count(&totalCount)

		database.DB.Where("id_category = ?", categoryID).
			Offset(offset).
			Limit(limit).
			Find(&products)

		totalPages := int((totalCount + int64(limit) - 1) / int64(limit))
		hasNext := page < totalPages
		hasPrev := page > 1

		c.HTML(200, "category.tmpl", gin.H{
			"products":    products,
			"currentPage": page,
			"totalPages":  totalPages,
			"hasNext":     hasNext,
			"hasPrev":     hasPrev,
			"categoryID":  categoryID,
			"totalCount":  totalCount,
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

	r.GET("/products/:id", func(c *gin.Context) {
		var product models.Product
		var categories []models.Category
		database.DB.Find(&categories)

		if err := database.DB.First(&product, c.Param("id")).Error; err != nil {
			c.String(404, "Product not found")
			return
		}

		c.HTML(200, "updateProduct.tmpl", gin.H{
			"product":    product,
			"categories": categories,
		})
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		if err := database.DB.Delete(&models.Product{}, c.Param("id")).Error; err != nil {
			c.String(500, "Internal server error")
			return
		}
		c.String(200, "Product deleted")
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		var product models.Product
		if err := database.DB.First(&product, c.Param("id")).Error; err != nil {
			log.Println("Product not found:", err)
			c.String(404, "Product not found")
			return
		}

		if err := c.ShouldBindJSON(&product); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(400, gin.H{"error": "Invalid JSON data: " + err.Error()})
			return
		}

		if err := database.DB.Save(&product).Error; err != nil {
			log.Println("Error saving product:", err)
			c.String(500, "Internal server error")
			return
		}

		log.Printf("Product updated successfully: %+v", product)
		c.JSON(200, gin.H{"message": "Product updated successfully"})
	})

	r.Run(":3000")
}

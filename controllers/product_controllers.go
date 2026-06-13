package controllers

import (
	"go-restfulapi/config"
	"go-restfulapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context){
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"data" : products,
		"total" : len(products),
	})
}

func GetProductById(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status" : "error",
			"message" : "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data" : product,
	})
}

func CreateProduct(c *gin.Context){
	var input struct{
		Name string `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"required"`
		Stock int `json:"stock" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "error",
			"message" : err.Error(),
		})
		return
	}

	product := models.Product{
		Name: input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, gin.H{
		"status" : "success",
		"message" : "Product created successfully",
		"data" : product,
	})
}

func UpdateProduct(c *gin.Context){
	var product models.Product
	id := c.Param("id")
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status" : "error",
			"message" : "Product not found",
		})
		return
	}

	var input struct{
		Name string `json:"name"`
		Price float64 `json:"price"`
		Stock int `json:"stock"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "error",
			"message" : err.Error(),
		})
		return
	}

	if input.Name != "" {
		product.Name = input.Name
	}

	if input.Price != 0 {
		product.Price = input.Price
	}

	if input.Stock != 0 {
		product.Stock = input.Stock
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Data updated successfully",
		"data": product, 
	})
}

func DeleteProductById(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "Product not found",
		})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Product deleted successfully",
	})
}
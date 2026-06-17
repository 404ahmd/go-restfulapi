package controllers

import (
	"go-restfulapi/config"
	"go-restfulapi/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil{
		config.InternalError(c, "Failed while get customer data", err.Error())
		return
	}
	config.OK(c, "Semua data customer berhasil diambil", products)
}

func GetProductById(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		config.NotFound(c, "Product Not Found", err.Error())
		return
	}

	config.OK(c, "Product Found", product)
}

func CreateProduct(c *gin.Context) {
	var input struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"required"`
		Stock int     `json:"stock" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		config.BadRequest(c, "Parameter Not Accepted", err.Error())
		return
	}

	product := models.Product{
		Name:  input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	config.DB.Create(&product)
	config.OK(c, "Product Created Successfully", product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.First(&product, id).Error; err != nil {
		config.NotFound(c, "Product Not Found", err.Error())
		return
	}

	var input struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		config.BadRequest(c, "Paramater Not Accepted", err.Error())
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
	config.OK(c, "Product Updated Successfully", product)
}

func DeleteProductById(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		config.NotFound(c, "Product Not Found", err.Error())
		return
	}

	config.DB.Delete(&product)
	config.OK(c, "Product Deleted Successfully", product)
}

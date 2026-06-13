package controllers

import (
	"go-restfulapi/config"
	"go-restfulapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomer(c *gin.Context){
	var customers []models.Customer
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"data" : customers,
		"total" : len(customers),
	})
}

func GetCustomerById(c *gin.Context){
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.Find(&customer, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "Customer not found",
		})
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"status":"success",
		"data": customer,
	})
}

func CreateCustomer(c *gin.Context){
	var input struct{
		Username string `json:"username" binding:"required"`
		Fullname string `json:"fullname" binding:"required"`
		Email string `json:"email" binding:"required"`
		Age int64 `json:"age" binding:"required"`
		Address string `json:"address" binding:"required"`
		Gender string `json:"gender" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "error",
			"message" : err.Error(),
		})
		return
	}

	customer := models.Customer{
		Username: input.Username,
		Fullname: input.Fullname,
		Email: input.Email,
		Age: input.Age,
		Address: input.Address,
		Gender: input.Gender,
	}

	config.DB.Create(&customer)
	c.JSON(http.StatusCreated, gin.H{
		"status" : "success",
		"message" : "Customer created successfully",
		"data" : customer,
	})
}

func UpdateCustomer(c *gin.Context){
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "Customer not found",
		})
		return
	}

	var input struct{
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Email string `json:"email"`
		Age int64 `json:"age"`
		Address string `json:"address"`
		Gender string `json:"gender"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":"error",
			"message":err.Error(),
		})
		return
	}

	if input.Username != "" {
		customer.Username = input.Username
	}

	if input.Fullname != "" {
		customer.Fullname = input.Fullname
	}

	if input.Email != "" {
		customer.Email = input.Email
	}

	if input.Age != 0 {
		customer.Age = input.Age
	}

	if input.Address != "" {
		customer.Address = input.Address
	}

	if input.Gender != "" {
		customer.Gender = input.Gender
	}

	config.DB.Save(&customer)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Customer data successfully updated",
		"data": customer,
	})

}

func DeleteCustomerById(c *gin.Context){
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "Product not found",
		})
		return
	}

	config.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Customer deleted successfully",
	})
}


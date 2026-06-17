package controllers

import (
	"go-restfulapi/config"
	"go-restfulapi/models"

	"github.com/gin-gonic/gin"
)

func GetCustomer(c *gin.Context) {
	var customers []models.Customer
	if err := config.DB.Find(&customers).Error; err != nil{
		config.InternalError(c, "Terjadi kesalahan di server", err.Error())
		return
	}
	config.OK(c, "Semua data customer berhasil diambil", customers)
}

func GetCustomerById(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		config.NotFound(c, "Customer not found", err.Error())
		return
	}

	config.OK(c, "Customer found", customer)
}

func CreateCustomer(c *gin.Context) {
	var input models.CreateCustomerRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		config.BadRequest(c, "Parameter not accepted", err.Error())
		return
	}

	var existingCustomer models.Customer
	if err := config.DB.Where("username = ? ", input.Username).First(&existingCustomer).Error; err != nil {
		config.BadRequest(c, "Username Already Registered", err.Error())
		return
	}

	if err := config.DB.Where("email = ? ", input.Email).First(&existingCustomer).Error; err!=nil {
		config.BadRequest(c, "Email Already Registered", err.Error())
		return
	}

	customer := models.Customer{
		Username: input.Username,
		Fullname: input.Fullname,
		Email:    input.Email,
		Age:      input.Age,
		Address:  input.Address,
		Gender:   input.Gender,
	}

	config.DB.Create(&customer)
	config.Created(c, "Customer created successfully", customer)
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		config.NotFound(c, "Customer not found", err.Error())
		return
	}

	var input models.UpdateCustomerRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		config.BadRequest(c, "Parameter not accepted", err.Error())
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
	config.OK(c, "Customer updated successfully", customer)

}

func DeleteCustomerById(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		config.NotFound(c, "Customer Not Found", err.Error())
		return
	}

	config.DB.Delete(&customer)
	config.OK(c, "Customer successfully deleted", customer)
}

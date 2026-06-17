package controllers

import (
	"go-restfulapi/config"
	"go-restfulapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string `binding:"required"`
		Email    string `binding:"required"`
		Password string `binding:"required"`
		Role     string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var existing models.User
	if result := config.DB.Where("Email = ?", input.Email).First(&existing); result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "error",
			"message": "Email already registered",
		})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to generate password",
		})
		return
	}

	inputRole := input.Role
	if inputRole == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "role cant be empty",
		})
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		Role:     inputRole,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	token, err := config.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, AuthResponse{Token: token, User: user})
}

func Login(c *gin.Context) {
	var Input struct {
		Email    string `binding:"required"`
		Password string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ? ", Input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token, err := config.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate password",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{Token: token, User: user})

}

func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var user models.User
	config.DB.First(&user, userID)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data user already",
		"data":    user,
	})
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

package config

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func OK(c *gin.Context, message string, data interface{}){
	c.JSON(200, Response{
		Success: true,
		Message: message,
		Data: data,
	})
}

func Created(c *gin.Context, message string, data interface{}){
	c.JSON(201, Response{
		Success: true,
		Message: message,
		Data: data,
	})
}

func BadRequest(c *gin.Context, message string, error interface{}){
	c.JSON(400, Response{
		Success: false,
		Message: message,
		Error: error,
	})
}

func Unauthorized(c *gin.Context, message string, error interface{}){
	c.JSON(401, Response{
		Success: false,
		Message: message,
		Error: error,
	})
}

func InternalError(c *gin.Context, message string, error interface{}){
	c.JSON(500, Response{
		Success: false,
		Message: message,
		Error: error,
	})
}

func NotFound(c *gin.Context, message string, error interface{}){
	c.JSON(404, Response{
		Success: false,
		Message: message,
		Error: error,
	})
}





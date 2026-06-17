package config

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte{
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "M3sSgzb5TdMpsefqSpLYo9zwAiAlx5gJLzYTGLt9utM"
	}

	return []byte(secret)
}

func GenerateToken(UserID uint, email, role string) (string, error) {
	claims := Claims{
		UserID: UserID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "go-restfulapi",
		},
	}
	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJWTSecret())
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		AuthHeader := ctx.GetHeader("Authorization")
		if AuthHeader == "" || !strings.HasPrefix(AuthHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization is Invalid",
			})
			return
		}

		tokenStr := strings.TrimPrefix(AuthHeader, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(GetJWTSecret()), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token" + err.Error(),
			})
			return
		}

		// Set user info ke context
		ctx.Set("user_id", claims.UserID)
		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}

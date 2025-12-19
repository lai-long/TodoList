package middleware

import (
	"TodoList/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthConfirm(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "token is needed"})
		c.Abort()
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer" {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.JwtSecret, nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "login err"})
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Set("username", claims["username"])
	c.Next()
}
func GetUsername(c *gin.Context) string {
	username := c.GetString("username")
	return username
}

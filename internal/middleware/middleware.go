package middleware

import (
	"TodoList/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// 确认是否成功登录的中间件
func AuthConfirm(c *gin.Context) {
	//获取密钥
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "token is needed"})
		c.Abort()
		return
	}
	//删除token前面的Bearer
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	//验证token是否正确，若正确返回密钥
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.JwtSecret, nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "login err"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "login err token"})
		c.Abort()
		return
	} else {
		c.Set("username", claims["username"])
		c.Next()
	}
}
func GetUsername(c *gin.Context) string {
	username := c.GetString("username")
	return username
}

package controlller

import (
	"TodoList/config"
	"TodoList/internal/model/dto"
	"TodoList/internal/model/entity"
	"TodoList/pkg/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var User entity.User

func Register(c *gin.Context) {
	var req dto.UserInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "register err",
			"error": err.Error()})
	}
	var count int64
	database.DB.Model(&entity.User{}).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户已存在"})
		return
	}
	User.UserName = req.UserName
	User.Password = req.Password
	User.Email = req.Email
	User.Phone = req.Phone
	User.Address = req.Address
	database.DB.Create(&User)
	c.JSON(http.StatusOK, gin.H{"msg": "register success"})
}
func Login(c *gin.Context) {
	var req dto.UserInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "login err",
			"error": err.Error()})
	}
	var user entity.User
	result := database.DB.Where("username = ? AND password =?", req.UserName, req.Password).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名不存在或密码错误"})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": User.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "login err"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "login success",
		"token": tokenString})
}

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

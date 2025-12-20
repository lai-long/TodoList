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

var Token string
var User entity.User

// 注册
func Register(c *gin.Context) {
	var req dto.UserInfo
	//获取前端数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "register err",
			"error": err.Error()})
		return
	}
	//确认用户是否重复
	var count int64
	database.DB.Model(&entity.User{}).Where("user_name=?", req.UserName).Count(&count)
	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户已存在",
			"count": count})
		return
	}
	//将dto转化为entity
	User.UserName = req.UserName
	User.Password = req.Password
	User.Email = req.Email
	User.Phone = req.Phone
	User.Address = req.Address
	//保存数据并返回给前端
	database.DB.Create(&User)
	c.JSON(http.StatusOK, gin.H{"msg": "register success"})
}

// 登录
func Login(c *gin.Context) {
	//获取前端数据
	var req dto.UserInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "login err bind-err",
			"error": err.Error()})
		return
	}
	var user entity.User
	//寻找用户在数据库中的信息
	result := database.DB.Where("user_name = ? AND password =?", req.UserName, req.Password).First(&user)
	//确认用户是否存在或密码是否正确
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名不存在或密码错误"})
		return
	}
	//获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": User.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	//用密钥获取完整的jwt字符串
	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "login err token-err",
			"error": err.Error()})
		return
	}
	//传输完整token给中间件或前端
	c.JSON(http.StatusOK, gin.H{"msg": "login success",
		"token": tokenString})
	Token = tokenString
}

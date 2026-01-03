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
	"golang.org/x/crypto/bcrypt"
)

var User entity.User

// Register godoc
// @Summary      用户注册
// @Description  新用户注册接口
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        request  body  dto.UserInfo  true  "用户注册信息"
// @Success      200  {object}  map[string]interface{}  "注册成功"
// @Failure      400  {object}  map[string]interface{}  "注册失败"
// @Router       /user/register [post]
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
	//加密密码
	hasherPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hashedPassword err",
		})
	}
	//将dto转化为entity
	User.UserName = req.UserName
	User.Password = string(hasherPassword)
	User.Email = req.Email
	User.Phone = req.Phone
	User.Address = req.Address
	//保存数据并返回给前端
	database.DB.Create(&User)
	c.JSON(http.StatusOK, gin.H{"msg": "register success"})
}

// Login godoc
// @Summary      用户登录
// @Description  用户登录接口，返回JWT令牌
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        request  body  dto.UserInfo  true  "用户登录信息"
// @Success      200  {object}  map[string]interface{}  "登录成功，返回token"
// @Failure      400  {object}  map[string]interface{}  "登录失败"
// @Router       /user/login [post]
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
	database.DB.Where("user_name = ?", req.UserName).First(&user)
	//确认用户是否存在或密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
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
}

package config

// 数据库连接字符串
var Dsn = "root:root@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"

// JWT密钥
var JwtSecret = []byte("secret-key")

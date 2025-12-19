package dto

type UserInfo struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Address  string `json:"address"`
}

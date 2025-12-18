package dto

type UserInfo struct {
	ID       int
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Address  string `json:"address"`
}

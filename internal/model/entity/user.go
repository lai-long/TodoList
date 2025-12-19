package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string `gorm:"not null VARCHAR(64)"`
	Email    string `gorm:"default:'未知'"`
	Phone    string `gorm:"default:'未知'"`
	Sex      string `gorm:"default:'沃尔玛购物袋'"`
	Address  string `gorm:"default:'未知'"`
}

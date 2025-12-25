package entity

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	UserName string
	Title    string `gorm:"size:255;default:'未命名'"`
	Context  string `gorm:"default:'未输入文本'"`
	Status   string `gorm:"default:'未完成'"`
	EndDate  string `gorm:"default:'2006-01-02 15:04:05'"`
}

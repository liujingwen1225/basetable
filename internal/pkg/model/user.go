package model

import "gorm.io/gorm"

// UserM 是数据库中 user 记录 struct 格式的映射.
type UserM struct {
	gorm.Model
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Nickname string `gorm:"column:nickname"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`
	Gender   int8   `gorm:"column:gender"`
}

// TableName 用来指定映射的 MySQL 表名.
func (u *UserM) TableName() string {
	return "user"
}

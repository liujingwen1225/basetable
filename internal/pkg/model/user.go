package model

const TableNameUserM = "user"

// UserM mapped from table <user>
type UserM struct {
	BaseModelM
	Username string `gorm:"column:username;type:longtext;not null" json:"username"`
	Password string `gorm:"column:password;type:longtext;not null" json:"password"`
	Nickname string `gorm:"column:nickname;type:longtext" json:"nickname"`
	Email    string `gorm:"column:email;type:longtext" json:"email"`
	Phone    string `gorm:"column:phone;type:longtext" json:"phone"`
	Gender   int    `gorm:"column:gender;type:tinyint(4)" json:"gender"`
}

// TableName UserM's table name
func (*UserM) TableName() string {
	return TableNameUserM
}

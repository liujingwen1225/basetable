// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePostM = "post"

// PostM mapped from table <post>
type PostM struct {
	ID        int       `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Username  string    `gorm:"column:username;type:varchar(255);not null" json:"username"`
	PostID    string    `gorm:"column:postID;type:varchar(256);not null" json:"postID"`
	Title     string    `gorm:"column:title;type:varchar(256);not null" json:"title"`
	Content   string    `gorm:"column:content;type:longtext;not null" json:"content"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

// TableName PostM's table name
func (*PostM) TableName() string {
	return TableNamePostM
}

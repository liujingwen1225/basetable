// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCollectionsM = "collections"

// CollectionsM mapped from table <collections>
type CollectionsM struct {
	ID          int            `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"`
	Name        string         `gorm:"column:name;type:varchar(100);not null;comment:名称" json:"name"`
	Type        int            `gorm:"column:type;type:tinyint(4);not null;comment:类型" json:"type"`
	SourceTable string         `gorm:"column:source_table;type:varchar(100);not null;comment:源表名称" json:"source_table"`
	Options     string         `gorm:"column:options;type:varchar(100);comment:额外参数" json:"options"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at"`
}

// TableName CollectionsM's table name
func (*CollectionsM) TableName() string {
	return TableNameCollectionsM
}

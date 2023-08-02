package model

import "gorm.io/gorm"

type SchemaM struct {
	gorm.Model
	Name          string  `gorm:"column:name;type:varchar(100);not null;comment:字段名称" json:"name"`
	Type          string  `gorm:"column:type;type:varchar(100);not null;comment:字段类型" json:"type"`
	Required      int     `gorm:"column:required;type:tinyint(1);not null;comment:是否必填" json:"required"`
	System        int     `gorm:"column:system;type:tinyint(1);not null;comment:系统字段" json:"system"`
	Options       *string `gorm:"column:options;type:varchar(100);comment:额外参数" json:"options"`
	CollectionsId int     `gorm:"column:collections_id;type:bigint;not null;comment:集合id" json:"collectionsId"`
}

// TableName CollectionsFieldsM's table name
func (*SchemaM) TableName() string {
	return "schema"
}

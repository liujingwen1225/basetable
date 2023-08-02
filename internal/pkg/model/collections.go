package model

import "gorm.io/gorm"

type CollectionsM struct {
	gorm.Model
	Name        string  `gorm:"column:name;type:varchar(100);not null;comment:名称" json:"name"`
	Type        int     `gorm:"column:type;type:tinyint(4);not null;comment:类型" json:"type"`
	SourceTable string  `gorm:"column:source_table;type:varchar(100);not null;comment:源表名称" json:"source_table"`
	Options     *string `gorm:"column:options;type:varchar(100);comment:额外参数" json:"options"`
}

func (*CollectionsM) TableName() string {
	return "collections"
}

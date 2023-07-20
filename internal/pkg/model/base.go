package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModelM struct {
	ID        int            `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"-"`
}

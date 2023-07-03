package db

import (
	"basetable.com/pkg/api"
	"gorm.io/gorm"
)

// 分页封装
func Paginate(pageM *api.PageRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageM.Page == 0 {
			pageM.PageSize = 10
		}
		switch {
		case pageM.PageSize > 500:
			pageM.PageSize = 500
		case pageM.PageSize <= 0:
			pageM.PageSize = 10
		}
		offset := (pageM.Page - 1) * pageM.PageSize
		return db.Offset(offset).Limit(pageM.PageSize)
	}
}

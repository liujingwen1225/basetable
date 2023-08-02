package store

import (
	"basetable.com/internal/pkg/model"
	v1 "basetable.com/pkg/api/basetable/v1"
	"context"
	"gorm.io/gorm"
)

type CollectionsStore interface {
	Create(cxt context.Context, collections *model.CollectionsM) error
	List(cxt context.Context, collections *model.CollectionsM, pagination v1.Pagination) ([]*model.CollectionsM, error)
	GetById(cxt context.Context, id int) (*model.CollectionsM, error)
	Update(cxt context.Context, collections *model.CollectionsM) error
	Delete(cxt context.Context, ids []int) error
}

var _ CollectionsStore = (*collections)(nil)

type collections struct {
	db *gorm.DB
}

func (c *collections) Create(cxt context.Context, collections *model.CollectionsM) error {
	return c.db.Save(&collections).Error
}

func (c *collections) List(cxt context.Context, collections *model.CollectionsM, pagination v1.Pagination) ([]*model.CollectionsM, error) {
	var res []*model.CollectionsM
	offset, limit := pagination.GetPage()
	query := c.db
	if collections.Name != "" {
		query = query.Where("name like ?", "%"+collections.Name+"%")
	}
	tx := query.Find(&res).Offset(offset).Limit(limit)
	return res, tx.Error
}

func (c *collections) GetById(cxt context.Context, id int) (*model.CollectionsM, error) {
	var res model.CollectionsM
	err := c.db.Where("id = ?", id).First(&res).Error
	return &res, err
}

func (c *collections) Update(cxt context.Context, collections *model.CollectionsM) error {
	return c.db.Updates(&collections).Error
}

func (c *collections) Delete(cxt context.Context, ids []int) error {
	return c.db.Delete(&model.CollectionsM{}, ids).Error
}

func newCollectionsStore(db *gorm.DB) *collections {
	return &collections{db: db}
}

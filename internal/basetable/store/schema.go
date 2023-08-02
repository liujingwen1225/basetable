package store

import (
	"basetable.com/internal/pkg/model"
	v1 "basetable.com/pkg/api/basetable/v1"
	"context"
	"gorm.io/gorm"
)

type SchemaStore interface {
	Create(cxt context.Context, schema *model.SchemaM) error
	List(cxt context.Context, schema *model.SchemaM, pagination v1.Pagination) ([]*model.SchemaM, error)
	GetById(cxt context.Context, id int) (*model.SchemaM, error)
	Update(cxt context.Context, schema *model.SchemaM) error
	Delete(cxt context.Context, ids []int) error
}

var _ SchemaStore = (*schema)(nil)

type schema struct {
	db *gorm.DB
}

func (c *schema) Create(cxt context.Context, schema *model.SchemaM) error {
	return c.db.Save(&schema).Error
}

func (c *schema) List(cxt context.Context, schema *model.SchemaM, pagination v1.Pagination) ([]*model.SchemaM, error) {
	var res []*model.SchemaM
	offset, limit := pagination.GetPage()
	query := c.db
	if schema.Name != "" {
		query = query.Where("name like ?", "%"+schema.Name+"%")
	}
	tx := query.Find(&res).Offset(offset).Limit(limit)
	return res, tx.Error
}

func (c *schema) GetById(cxt context.Context, id int) (*model.SchemaM, error) {
	var res model.SchemaM
	err := c.db.Where("id = ?", id).First(&res).Error
	return &res, err
}

func (c *schema) Update(cxt context.Context, schema *model.SchemaM) error {
	return c.db.Updates(&schema).Error
}

func (c *schema) Delete(cxt context.Context, ids []int) error {
	return c.db.Delete(&model.SchemaM{}, ids).Error
}

func newSchemaStore(db *gorm.DB) *schema {
	return &schema{db: db}
}

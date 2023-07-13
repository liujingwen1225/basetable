// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package store

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"basetable.com/internal/pkg/model"
)

func newCollectionsFieldsM(db *gorm.DB, opts ...gen.DOOption) collectionsFieldsM {
	_collectionsFieldsM := collectionsFieldsM{}

	_collectionsFieldsM.collectionsFieldsMDo.UseDB(db, opts...)
	_collectionsFieldsM.collectionsFieldsMDo.UseModel(&model.CollectionsFieldsM{})

	tableName := _collectionsFieldsM.collectionsFieldsMDo.TableName()
	_collectionsFieldsM.ALL = field.NewAsterisk(tableName)
	_collectionsFieldsM.ID = field.NewInt(tableName, "id")
	_collectionsFieldsM.CollectionsID = field.NewInt(tableName, "collections_id")
	_collectionsFieldsM.Name = field.NewString(tableName, "name")
	_collectionsFieldsM.Type = field.NewString(tableName, "type")
	_collectionsFieldsM.Required = field.NewInt(tableName, "required")
	_collectionsFieldsM.System = field.NewInt(tableName, "system")
	_collectionsFieldsM.Options = field.NewString(tableName, "options")
	_collectionsFieldsM.CreatedAt = field.NewTime(tableName, "created_at")
	_collectionsFieldsM.UpdatedAt = field.NewTime(tableName, "updated_at")
	_collectionsFieldsM.DeletedAt = field.NewField(tableName, "deleted_at")

	_collectionsFieldsM.fillFieldMap()

	return _collectionsFieldsM
}

type collectionsFieldsM struct {
	collectionsFieldsMDo

	ALL           field.Asterisk
	ID            field.Int
	CollectionsID field.Int    // 集合id
	Name          field.String // 字段名称
	Type          field.String // 字段类型
	Required      field.Int    // 是否必填
	System        field.Int    // 系统字段
	Options       field.String // 额外参数
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (c collectionsFieldsM) Table(newTableName string) *collectionsFieldsM {
	c.collectionsFieldsMDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c collectionsFieldsM) As(alias string) *collectionsFieldsM {
	c.collectionsFieldsMDo.DO = *(c.collectionsFieldsMDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *collectionsFieldsM) updateTableName(table string) *collectionsFieldsM {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt(table, "id")
	c.CollectionsID = field.NewInt(table, "collections_id")
	c.Name = field.NewString(table, "name")
	c.Type = field.NewString(table, "type")
	c.Required = field.NewInt(table, "required")
	c.System = field.NewInt(table, "system")
	c.Options = field.NewString(table, "options")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *collectionsFieldsM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *collectionsFieldsM) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["collections_id"] = c.CollectionsID
	c.fieldMap["name"] = c.Name
	c.fieldMap["type"] = c.Type
	c.fieldMap["required"] = c.Required
	c.fieldMap["system"] = c.System
	c.fieldMap["options"] = c.Options
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c collectionsFieldsM) clone(db *gorm.DB) collectionsFieldsM {
	c.collectionsFieldsMDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c collectionsFieldsM) replaceDB(db *gorm.DB) collectionsFieldsM {
	c.collectionsFieldsMDo.ReplaceDB(db)
	return c
}

type collectionsFieldsMDo struct{ gen.DO }

type ICollectionsFieldsMDo interface {
	gen.SubQuery
	Debug() ICollectionsFieldsMDo
	WithContext(ctx context.Context) ICollectionsFieldsMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICollectionsFieldsMDo
	WriteDB() ICollectionsFieldsMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICollectionsFieldsMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICollectionsFieldsMDo
	Not(conds ...gen.Condition) ICollectionsFieldsMDo
	Or(conds ...gen.Condition) ICollectionsFieldsMDo
	Select(conds ...field.Expr) ICollectionsFieldsMDo
	Where(conds ...gen.Condition) ICollectionsFieldsMDo
	Order(conds ...field.Expr) ICollectionsFieldsMDo
	Distinct(cols ...field.Expr) ICollectionsFieldsMDo
	Omit(cols ...field.Expr) ICollectionsFieldsMDo
	Join(table schema.Tabler, on ...field.Expr) ICollectionsFieldsMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICollectionsFieldsMDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICollectionsFieldsMDo
	Group(cols ...field.Expr) ICollectionsFieldsMDo
	Having(conds ...gen.Condition) ICollectionsFieldsMDo
	Limit(limit int) ICollectionsFieldsMDo
	Offset(offset int) ICollectionsFieldsMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectionsFieldsMDo
	Unscoped() ICollectionsFieldsMDo
	Create(values ...*model.CollectionsFieldsM) error
	CreateInBatches(values []*model.CollectionsFieldsM, batchSize int) error
	Save(values ...*model.CollectionsFieldsM) error
	First() (*model.CollectionsFieldsM, error)
	Take() (*model.CollectionsFieldsM, error)
	Last() (*model.CollectionsFieldsM, error)
	Find() ([]*model.CollectionsFieldsM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CollectionsFieldsM, err error)
	FindInBatches(result *[]*model.CollectionsFieldsM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CollectionsFieldsM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICollectionsFieldsMDo
	Assign(attrs ...field.AssignExpr) ICollectionsFieldsMDo
	Joins(fields ...field.RelationField) ICollectionsFieldsMDo
	Preload(fields ...field.RelationField) ICollectionsFieldsMDo
	FirstOrInit() (*model.CollectionsFieldsM, error)
	FirstOrCreate() (*model.CollectionsFieldsM, error)
	FindByPage(offset int, limit int) (result []*model.CollectionsFieldsM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICollectionsFieldsMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c collectionsFieldsMDo) Debug() ICollectionsFieldsMDo {
	return c.withDO(c.DO.Debug())
}

func (c collectionsFieldsMDo) WithContext(ctx context.Context) ICollectionsFieldsMDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c collectionsFieldsMDo) ReadDB() ICollectionsFieldsMDo {
	return c.Clauses(dbresolver.Read)
}

func (c collectionsFieldsMDo) WriteDB() ICollectionsFieldsMDo {
	return c.Clauses(dbresolver.Write)
}

func (c collectionsFieldsMDo) Session(config *gorm.Session) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Session(config))
}

func (c collectionsFieldsMDo) Clauses(conds ...clause.Expression) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c collectionsFieldsMDo) Returning(value interface{}, columns ...string) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c collectionsFieldsMDo) Not(conds ...gen.Condition) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c collectionsFieldsMDo) Or(conds ...gen.Condition) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c collectionsFieldsMDo) Select(conds ...field.Expr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c collectionsFieldsMDo) Where(conds ...gen.Condition) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c collectionsFieldsMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICollectionsFieldsMDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c collectionsFieldsMDo) Order(conds ...field.Expr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c collectionsFieldsMDo) Distinct(cols ...field.Expr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c collectionsFieldsMDo) Omit(cols ...field.Expr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c collectionsFieldsMDo) Join(table schema.Tabler, on ...field.Expr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c collectionsFieldsMDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c collectionsFieldsMDo) RightJoin(table schema.Tabler, on ...field.Expr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c collectionsFieldsMDo) Group(cols ...field.Expr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c collectionsFieldsMDo) Having(conds ...gen.Condition) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c collectionsFieldsMDo) Limit(limit int) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c collectionsFieldsMDo) Offset(offset int) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c collectionsFieldsMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c collectionsFieldsMDo) Unscoped() ICollectionsFieldsMDo {
	return c.withDO(c.DO.Unscoped())
}

func (c collectionsFieldsMDo) Create(values ...*model.CollectionsFieldsM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c collectionsFieldsMDo) CreateInBatches(values []*model.CollectionsFieldsM, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c collectionsFieldsMDo) Save(values ...*model.CollectionsFieldsM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c collectionsFieldsMDo) First() (*model.CollectionsFieldsM, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionsFieldsM), nil
	}
}

func (c collectionsFieldsMDo) Take() (*model.CollectionsFieldsM, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionsFieldsM), nil
	}
}

func (c collectionsFieldsMDo) Last() (*model.CollectionsFieldsM, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionsFieldsM), nil
	}
}

func (c collectionsFieldsMDo) Find() ([]*model.CollectionsFieldsM, error) {
	result, err := c.DO.Find()
	return result.([]*model.CollectionsFieldsM), err
}

func (c collectionsFieldsMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CollectionsFieldsM, err error) {
	buf := make([]*model.CollectionsFieldsM, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c collectionsFieldsMDo) FindInBatches(result *[]*model.CollectionsFieldsM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c collectionsFieldsMDo) Attrs(attrs ...field.AssignExpr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c collectionsFieldsMDo) Assign(attrs ...field.AssignExpr) ICollectionsFieldsMDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c collectionsFieldsMDo) Joins(fields ...field.RelationField) ICollectionsFieldsMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c collectionsFieldsMDo) Preload(fields ...field.RelationField) ICollectionsFieldsMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c collectionsFieldsMDo) FirstOrInit() (*model.CollectionsFieldsM, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionsFieldsM), nil
	}
}

func (c collectionsFieldsMDo) FirstOrCreate() (*model.CollectionsFieldsM, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionsFieldsM), nil
	}
}

func (c collectionsFieldsMDo) FindByPage(offset int, limit int) (result []*model.CollectionsFieldsM, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c collectionsFieldsMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c collectionsFieldsMDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c collectionsFieldsMDo) Delete(models ...*model.CollectionsFieldsM) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *collectionsFieldsMDo) withDO(do gen.Dao) *collectionsFieldsMDo {
	c.DO = *do.(*gen.DO)
	return c
}

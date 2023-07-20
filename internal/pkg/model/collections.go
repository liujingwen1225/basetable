package model

const TableNameCollectionsM = "collections"

// CollectionsM mapped from table <collections>
type CollectionsM struct {
	BaseModelM
	Name        string                `gorm:"column:name;type:varchar(100);not null;comment:名称" json:"name"`
	Type        string                `gorm:"column:type;type:varchar(100);not null;comment:类型" json:"type"`
	SourceTable string                `gorm:"column:source_table;type:varchar(100);not null;comment:源表名称" json:"source_table"`
	Options     string                `gorm:"column:options;type:varchar(500);comment:额外参数" json:"options"`
	Fields      []*CollectionsFieldsM `gorm:"foreignKey:CollectionsID" json:"fields"`
}

// TableName CollectionsM's table name
func (*CollectionsM) TableName() string {
	return TableNameCollectionsM
}

const (
	CollectionTypeBase = "base"
	CollectionTypeAuth = "auth"
	CollectionTypeView = "view"
)

var CollectionTypes = []string{CollectionTypeBase, CollectionTypeAuth, CollectionTypeView}

func (c *CollectionsM) IsInCollectionTypes() bool {
	for _, v := range CollectionTypes {
		if v == c.Type {
			return true
		}
	}
	return false
}

func (c *CollectionsM) CollectionSystemFieldsInit() {
	switch c.Type {
	case CollectionTypeBase:
		c.Fields = append(c.Fields, BaseSystemFields()...)
	case CollectionTypeAuth:
	case CollectionTypeView:
	}
	//c.Fields = append()
}

func BaseSystemFields() []*CollectionsFieldsM {
	return []*CollectionsFieldsM{
		{Name: "id", Type: "number", Required: false, System: true, Options: "primaryKey"},
		{Name: "created", Type: "number", Required: false, System: true, Options: ""},
		{Name: "update", Type: "number", Required: false, System: true, Options: ""},
	}
}

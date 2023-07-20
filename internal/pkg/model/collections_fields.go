package model

const TableNameCollectionsFieldsM = "collections_fields"

// CollectionsFieldsM mapped from table <collections_fields>
type CollectionsFieldsM struct {
	BaseModelM
	Name          string `gorm:"column:name;type:varchar(100);not null;comment:字段名称" json:"name"`
	Type          string `gorm:"column:type;type:varchar(100);not null;comment:字段类型" json:"type"`
	Required      bool   `gorm:"column:required;type:bool;not null;comment:是否必填" json:"required"`
	System        bool   `gorm:"column:system;type:bool;not null;comment:系统字段" json:"system"`
	Options       string `gorm:"column:options;type:varchar(100);comment:额外参数" json:"options"`
	CollectionsID int    `gorm:"column:collections_id;type:bigint(20);not null;comment:集合id" json:"collections_id"`
}

// TableName CollectionsFieldsM's table name
func (*CollectionsFieldsM) TableName() string {
	return TableNameCollectionsFieldsM
}

// commonly used field names
const (
	FieldNameId                     string = "id"
	FieldNameCreated                string = "created"
	FieldNameUpdated                string = "updated"
	FieldNameCollectionId           string = "collectionId"
	FieldNameCollectionName         string = "collectionName"
	FieldNameExpand                 string = "expand"
	FieldNameUsername               string = "username"
	FieldNameEmail                  string = "email"
	FieldNameEmailVisibility        string = "emailVisibility"
	FieldNameVerified               string = "verified"
	FieldNameTokenKey               string = "tokenKey"
	FieldNamePasswordHash           string = "passwordHash"
	FieldNameLastResetSentAt        string = "lastResetSentAt"
	FieldNameLastVerificationSentAt string = "lastVerificationSentAt"
)

const (
	FieldTypeText     string = "text"
	FieldTypeNumber   string = "number"
	FieldTypeBool     string = "bool"
	FieldTypeEmail    string = "email"
	FieldTypeUrl      string = "url"
	FieldTypeEditor   string = "editor"
	FieldTypeDate     string = "date"
	FieldTypeSelect   string = "select"
	FieldTypeJson     string = "json"
	FieldTypeFile     string = "file"
	FieldTypeRelation string = "relation"

	// Deprecated: Will be removed in v0.9+
	FieldTypeUser string = "user"
)

var FieldTypes = []string{
	FieldTypeText,
	FieldTypeNumber,
	FieldTypeBool,
	FieldTypeEmail,
	FieldTypeUrl,
	FieldTypeEditor,
	FieldTypeDate,
	FieldTypeSelect,
	FieldTypeJson,
	FieldTypeFile,
	FieldTypeRelation,
}

func (f *CollectionsFieldsM) IsInFieldType() bool {
	for _, v := range FieldTypes {
		if v == f.Type {
			return true
		}
	}
	return false
}

// BaseModelFieldNames returns the field names that all models have (id, created, updated).
func BaseModelFieldNames() []string {
	return []string{
		FieldNameId,
		FieldNameCreated,
		FieldNameUpdated,
	}
}

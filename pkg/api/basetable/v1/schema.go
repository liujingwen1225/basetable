package v1

type SchemaRequest struct {
	Name     string  `valid:"required;alphanum;gt=30" json:"name"`
	Type     int     `valid:"required" json:"type"`
	Required *bool   `valid:"required" json:"required"`
	Options  *string `valid:"" json:"options"`
}

type SchemaResponse struct {
	Name     string  `gorm:"column:name;type:varchar(100);not null;comment:名称" json:"name"`
	Type     int     `gorm:"column:type;type:tinyint(4);not null;comment:类型" json:"type"`
	Required bool    `gorm:"column:options;type:varchar(100);comment:额外参数" json:"required"`
	Options  *string `gorm:"column:options;type:varchar(100);comment:额外参数" json:"options"`
}

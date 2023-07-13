package main

import (
	"basetable.com/internal/pkg/log"
	"basetable.com/pkg/db"
	"github.com/spf13/viper"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/basetable/store",
		OutFile:      "store.go",
		ModelPkgPath: "./internal/pkg/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: false, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	})
	g.WithModelNameStrategy(func(tableName string) (modelName string) {
		return camelString(tableName) + "M"
	})
	client, _ := db.NewDbClient(getDbOptions())
	g.UseDB(client) // reuse your gorm db
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	//gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//	toStringField := `banlance, `
	//	if strings.Contains(toStringField, columnName) {
	//		return columnName + ",string"
	//	}
	//	return columnName
	//})
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	//autoUpdateTimeField := gen.FieldGORMTag("update_time", func)
	//autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	//softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")
	//fieldOpts := []gen.ModelOpt{}
	//fieldOpts := []gen.ModelOpt{softDeleteField}

	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	//User := g.GenerateModel("user")
	// 创建全部模型文件, 并覆盖前面创建的同名模型
	allModel := g.GenerateAllTable()

	// 创建有关联关系的模型文件
	// 可以用于指定外键
	//Score := g.GenerateModel("score",
	// append(
	//    fieldOpts,
	//    // user 一对多 address 关联, 外键`uid`在 address 表中
	//    gen.FieldRelate(field.HasMany, "user", User, &field.RelateConfig{GORMTag: "foreignKey:UID"}),
	// )...,
	//)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(allModel...)

	g.Execute()
}

func getDbOptions() *db.MySqlOptions {

	viper.SetConfigFile("./configs/basetable.yaml")
	viper.AutomaticEnv()
	// 读取环境变量的前缀为 basetable，如果是 basetable，将自动转变为大写。
	viper.SetEnvPrefix("basetable.com")
	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}
	return &db.MySqlOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

/**
 * 蛇形转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY
 * @date 2020/7/30
 * @param s要转换的字符串
 * @return string
 **/
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

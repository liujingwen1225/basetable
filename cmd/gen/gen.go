package main

import (
	"basetable.com/internal/pkg/log"
	"basetable.com/internal/pkg/model"
	"basetable.com/pkg/db"
	"github.com/spf13/viper"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/basetable/store",
		OutFile: "store.go",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable
	})
	//g.WithModelNameStrategy(func(tableName string) (modelName string) {
	//	return camelString(tableName) + "M"
	//})
	client, _ := db.NewDbClient(getDbOptions())
	g.UseDB(client) // reuse your gorm db
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int" },
	}
	type Querier interface {
		// SELECT * FROM @@table WHERE LOWER(name)=@name
		CollectionNameByLowercase(name string) ([]*gen.T, error)
	}

	// 类型映射 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 需要迁移的模型
	allModel := []interface{}{
		model.UserM{},
		model.CollectionsM{},
		model.CollectionsFieldsM{},
	}
	// 自定义sql 查询
	g.ApplyInterface(func(Querier) {}, model.CollectionsM{})
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

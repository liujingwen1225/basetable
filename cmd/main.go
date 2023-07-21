package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Field struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Required bool        `json:"required"`
	Options  interface{} `json:"options"`
}

func createTableFromJSON(jsonData []byte) (string, error) {
	// 解析JSON数据，获取表名和字段列表
	var data struct {
		Name   string  `json:"name"`
		Fields []Field `json:"fields"`
	}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", err
	}

	// 构造创建表的SQL语句
	sqlStatement := fmt.Sprintf("CREATE TABLE %s (", data.Name)
	for i, field := range data.Fields {
		// 添加字段名称和类型
		sqlStatement += fmt.Sprintf("%s %s", field.Name, getSQLType(field.Type))

		// 如果字段为不可为空，则添加 NOT NULL 约束
		if field.Required {
			sqlStatement += " NOT NULL"
		}

		// 添加字段分隔符
		if i < len(data.Fields)-1 {
			sqlStatement += ", "
		}
	}
	sqlStatement += ");"

	return sqlStatement, nil
}

func getSQLType(fieldType string) string {
	// 根据字段类型返回对应的 MySQL 数据库类型
	switch fieldType {
	case "number":
		return "INT"
	case "text":
		return "TEXT"
	// 添加其他字段类型的映射关系
	default:
		return ""
	}
}

func main() {
	jsonData := []byte(`
		{
			"name": "my_table",
			"fields": [
				{
					"name": "age",
					"type": "number",
					"required": true,
					"options": {
						"min": 1,
						"max": 200
					}
				},
				{
					"name": "context",
					"type": "text",
					"required": false,
					"options": ""
				}
			]
		}
	`)

	sqlStatement, err := createTableFromJSON(jsonData)
	if err != nil {
		fmt.Println("Failed to generate SQL statement:", err)
		return
	}

	fmt.Println(sqlStatement)
}

func updateTableFromJSON(jsonData []byte) error {
	// 解析JSON数据，获取表名和字段列表
	var data struct {
		Name   string  `json:"name"`
		Fields []Field `json:"fields"`
	}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}

	// 获取数据库连接
	db, err := sql.Open("mysql", "username:password@tcp(hostname:port)/database_name")
	if err != nil {
		return err
	}
	defer db.Close()

	// 开始一个事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 检查表是否存在
	tableExists := false
	err = tx.QueryRow("SELECT 1 FROM information_schema.tables WHERE table_schema = ? AND table_name = ?", "database_name", data.Name).Scan(&tableExists)
	if err != nil {
		tx.Rollback()
		return err
	}

	if tableExists {
		// 获取当前表的字段列表
		currentFields, err := getTableFields(data.Name, tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 对比新的字段列表和当前表的字段列表，找出需要添加和删除的字段
		fieldsToAdd := findFieldsToAdd(currentFields, data.Fields)
		fieldsToDelete := findFieldsToDelete(currentFields, data.Fields)

		// 添加新的字段
		err = addFields(data.Name, fieldsToAdd, tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 删除不需要的字段
		err = deleteFields(data.Name, fieldsToDelete, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// 如果表不存在，直接创建新表
		err = createTableFromJSON(jsonData, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getTableFields(tableName string, tx *sql.Tx) ([]string, error) {
	// 查询表的字段列表
	var columns []string

	rows, err := tx.Query("SHOW COLUMNS FROM " + tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var column string
		err := rows.Scan(&column)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return columns, nil
}

func findFieldsToAdd(currentFields []string, newFields []Field) []Field {
	// 找出需要添加的字段
	var fieldsToAdd []Field

	currentFieldsMap := make(map[string]bool)
	for _, field := range currentFields {
		currentFieldsMap[field] = true
	}

	for _, field := range newFields {
		if !currentFieldsMap[field.Name] {
			fieldsToAdd = append(fieldsToAdd, field)
		}
	}

	return fieldsToAdd
}

func findFieldsToDelete(currentFields []string, newFields []Field) []string {
	// 找出需要删除的字段
	var fieldsToDelete []string

	newFieldsMap := make(map[string]bool)
	for _, field := range newFields {
		newFieldsMap[field.Name] = true
	}

	for _, field := range currentFields {
		if !newFieldsMap[field] {
			fieldsToDelete = append(fieldsToDelete, field)
		}
	}

	return fieldsToDelete
}

func addFields(tableName string, fields []Field, tx *sql.Tx) error {
	// 添加新的字段到表中
	for _, field := range fields {
		sqlStatement := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", tableName, field.Name, getSQLType(field.Type))
		if field.Required {
			sqlStatement += " NOT NULL"
		}

		_, err := tx.Exec(sqlStatement)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteFields(tableName string, fields []string, tx *sql.Tx) error {
	// 从表中删除不需要的字段
	for _, field := range fields {
		sqlStatement := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", tableName, field)

		_, err := tx.Exec(sqlStatement)
		if err != nil {
			return err
		}
	}

	return nil
}

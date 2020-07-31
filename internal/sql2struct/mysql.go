package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"

	// 引入数据库驱动注册及初始化
	_ "github.com/go-sql-driver/mysql"
)

// DBModel DB基础结构体
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

// DBInfo DB连接信息结构体
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// TableColumn COLUMN表结构体
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

// DBTypeToStructType 类型简单转换
var DBTypeToStructType = map[string]string{
	"int":       "int32",
	"tinyint":   "int8",
	"smallint":  "int",
	"mediumint": "int64",
	"bigint":    "int64",
	"bit":       "int",
	"bool":      "bool",
	"enum":      "string",
	"set":       "string",
	"varchar":   "string",
	"char":      "string",
	"text":      "string",
}

// NewDBModel 创建一个DB对象
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

// Connect 连接数据库
func (m *DBModel) Connect() error {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	// sql.Open() 第一参数为驱动名称，如mysql，第二个参数为驱动连接数据库的连接信息
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}

	return nil
}

// GetColumns 获取表中列表的信息
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " + "FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("no data")
	}
	defer rows.Close()

	var columns []*TableColumn

	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}

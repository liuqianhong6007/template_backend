package route

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"

	"github.com/liuqianhong6007/template_backend/config"
)

type Table struct {
	Name    string
	Comment string
	Columns []Column
}

type Column struct {
	ColumnName    string
	DataType      string
	ColumnType    string
	ColumnComment string
	ColumnKey     string
}

func GetTables(c *gin.Context) {
	rows, err := GetTablesStmt.QueryContext(c)
	checkValue(c, err)
	defer rows.Close()

	var tables []Table
	for rows.Next() {
		var table Table
		err = rows.Scan(&table.Name, &table.Comment)
		checkValue(c, err)
		tables = append(tables, table)
	}

	c.HTML(http.StatusOK, "getTables.tpl", tables)
}

func GetTable(c *gin.Context) {
	tableName := c.Query("tableName")
	rows, err := gDb.QueryContext(c, GetTableSql, config.DbCfg().LibName, tableName)
	checkValue(c, err)
	defer rows.Close()

	var columns []Column
	for rows.Next() {
		var column Column
		err = rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnType, &column.ColumnComment, &column.ColumnKey)
		checkValue(c, err)
		columns = append(columns, column)
	}
	table := Table{
		Name:    tableName,
		Columns: columns,
	}

	c.HTML(http.StatusOK, "getTable.tpl", table)
}

type TableRecord struct {
	TableName string
	Columns   []string
	Records   []Record
}

type Record struct {
	RecordColumns []RecordVal
}

type RecordVal struct {
	Editable   bool
	ColumnName string
	Val        string
}

func getTableMeta(c *gin.Context, tableName string) (columns []Column) {
	rows, err := gDb.QueryContext(c, GetTableSql, config.DbCfg().LibName, tableName)
	checkValue(c, err)
	defer rows.Close()

	for rows.Next() {
		var column Column
		err = rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnType, &column.ColumnComment, &column.ColumnKey)
		checkValue(c, err)
		columns = append(columns, column)
	}
	return
}

func getTableRecords(c *gin.Context, tableName string) {
	columns := getTableMeta(c, tableName)

	var tableRecord TableRecord
	if len(columns) == 0 {
		c.HTML(200, "getTableRecords.tpl", tableRecord)
		return
	}

	sqlStr := `select `
	realColumnVals := make([][]byte, len(columns))
	columnVals := make([]interface{}, len(columns))
	for i, column := range columns {
		sqlStr = sqlStr + column.ColumnName + `,`
		columnVals[i] = &realColumnVals[i]
		tableRecord.Columns = append(tableRecord.Columns, column.ColumnName)
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	sqlStr = sqlStr + ` from ` + tableName

	rowRecords, err := gDb.QueryContext(c, sqlStr)
	checkValue(c, err)
	defer rowRecords.Close()

	for rowRecords.Next() {

		err = rowRecords.Scan(columnVals...)
		checkValue(c, err)
		var record Record
		var recordVal RecordVal
		for i, realColumnVal := range realColumnVals {
			recordVal = RecordVal{
				ColumnName: columns[i].ColumnName,
				Val:        string(realColumnVal),
			}
			if columns[i].ColumnKey != "PRI" {
				recordVal.Editable = true
			}
			record.RecordColumns = append(record.RecordColumns, recordVal)
		}
		tableRecord.Records = append(tableRecord.Records, record)
	}
	tableRecord.TableName = tableName

	c.HTML(http.StatusOK, "getTableRecords.tpl", tableRecord)
}

func GetTableRecords(c *gin.Context) {
	tableName := c.Query("tableName")
	getTableRecords(c, tableName)
}

func UpdateTableRecord(c *gin.Context) {
	tableName := c.PostForm("tableName")
	checkValue(c, tableName != "", "param[tableName] is null")

	columns := getTableMeta(c, tableName)
	checkValue(c, len(columns) > 0, fmt.Sprintf("table[%s] not exit", tableName))

	columnMap := make(map[string]string)
	var pk string
	for _, column := range columns {
		if column.ColumnKey == "PRI" { // primary key
			pk = column.ColumnName
		}
		columnMap[column.ColumnName] = column.DataType
	}
	checkValue(c, pk != "", fmt.Sprintf("table[%s] has not pk yet", tableName))

	data := c.PostForm("data")
	checkValue(c, data != "", "data param is null")

	paramMap := make(map[string]string)
	var paramCnt int
	for _, kv := range strings.Split(data, ";") {
		arr := strings.Split(kv, "=")
		if len(arr) != 2 {
			continue
		}
		if _, ok := columnMap[arr[0]]; !ok {
			continue
		}
		paramMap[arr[0]] = arr[1]
		paramCnt++
	}
	if _, ok := paramMap[pk]; !ok {
		checkValue(c, errors.New("param must include pk"))
	}
	if paramCnt < 2 {
		checkValue(c, errors.New("param must include at least one column to update"))
	}
	sqlStr := "update " + tableName + " set "
	var args []interface{}
	for k, v := range paramMap {
		sqlStr = sqlStr + fmt.Sprintf("%s = ?,", k)
		args = append(args, transform(c, columnMap, k, v))
	}

	sqlStr = sqlStr[:len(sqlStr)-1]
	sqlStr = sqlStr + fmt.Sprintf(" where %s = ?", pk)
	args = append(args, paramMap[pk])

	_, err := gDb.ExecContext(c, sqlStr, args...)
	checkValue(c, err)

	getTableRecords(c, tableName)
}

func transform(c *gin.Context, columnMap map[string]string, k, v string) (val interface{}) {
	var err error
	switch columnMap[k] {
	case "tinyint", "smallint", "mediumint", "int", "integer", "bigint", "decimal", "numeric":
		val, err = strconv.Atoi(v)
		checkValue(c, err, "column data type mismatch")

	case "float", "double":
		val, err = strconv.ParseFloat(v, 64)
		checkValue(c, err, "column data type mismatch")

	case "char", "varchar", "tinytext", "mediumtext", "longtext", "text", "tinyblob", "mediumblob", "longblob", "blob":
		val = v
	default:
		checkValue(c, errors.New("unknown data type"))
	}
	return
}

func DeleteTableRecord(c *gin.Context) {
	tableName := c.PostForm("tableName")
	checkValue(c, tableName != "", "param[tableName] is null")

	columns := getTableMeta(c, tableName)
	checkValue(c, len(columns) > 0, fmt.Sprintf("table[%s] not exit", tableName))

	columnMap := make(map[string]string)
	var pk string
	for _, column := range columns {
		if column.ColumnKey == "PRI" { // primary key
			pk = column.ColumnName
		}
		columnMap[column.ColumnName] = column.DataType
	}
	checkValue(c, pk != "", fmt.Sprintf("table[%s] has not pk yet", tableName))

	data := c.PostForm("data")
	checkValue(c, data != "", "data param is null")

	paramMap := make(map[string]string)
	for _, kv := range strings.Split(data, ";") {
		arr := strings.Split(kv, "=")
		if len(arr) != 2 {
			continue
		}
		if _, ok := columnMap[arr[0]]; !ok {
			continue
		}
		paramMap[arr[0]] = arr[1]
	}
	if _, ok := paramMap[pk]; !ok {
		checkValue(c, errors.New("param must include pk"))
	}

	sqlStr := "delete from " + tableName + fmt.Sprintf(" where %s = ?", pk)
	var args []interface{}
	args = append(args, transform(c, columnMap, pk, paramMap[pk]))

	_, err := gDb.ExecContext(c, sqlStr, args...)
	checkValue(c, err)

	getTableRecords(c, tableName)
}

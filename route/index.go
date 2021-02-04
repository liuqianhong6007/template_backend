package route

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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

func Index(c *gin.Context) {
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

	c.HTML(200, "Index.tpl", tables)
}

func GetTable(c *gin.Context) {
	tableName := c.Query("tableName")
	rows, err := gDb.QueryContext(c, GetTableSql, config.DbCfg().LibName, tableName)
	checkValue(c, err)
	defer rows.Close()

	var columns []Column
	for rows.Next() {
		var column Column
		err = rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnType, &column.ColumnComment)
		checkValue(c, err)
		columns = append(columns, column)
	}
	table := Table{
		Name:    tableName,
		Columns: columns,
	}

	c.HTML(200, "GetTable.tpl", table)
}

type TableRecord struct {
	Columns []string
	Records []Record
}

type Record struct {
	Values []string
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

func GetTableRecords(c *gin.Context) {
	tableName := c.Query("tableName")
	columns := getTableMeta(c, tableName)

	var tableRecord TableRecord
	if len(columns) == 0 {
		c.HTML(200, "GetTableRecords.tpl", tableRecord)
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
		for _, realColumnVal := range realColumnVals {
			record.Values = append(record.Values, string(realColumnVal))
		}
		tableRecord.Records = append(tableRecord.Records, record)
	}

	c.HTML(200, "GetTableRecords.tpl", tableRecord)
}

func UpdateTableRecord(c *gin.Context) {
	tableName := c.PostForm("tableName")
	columns := getTableMeta(c, tableName)
	columnMap := make(map[string]string)
	var pk string
	for _, column := range columns {
		if column.ColumnKey == "PRI" { // primary key
			pk = column.ColumnName
		}
		columnMap[column.ColumnName] = column.ColumnName
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

}

package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/template_backend/config"
)

func init() {
	AddRoute(Route{
		Method:  http.MethodGet,
		Path:    "/index",
		Handler: index,
	})
	AddRoute(Route{
		Method:  http.MethodGet,
		Path:    "/getTable",
		Handler: getTable,
	})
	AddRoute(Route{
		Method:  http.MethodGet,
		Path:    "/getTableRecords",
		Handler: getTableRecords,
	})
}

func checkErr(c *gin.Context, err error) {
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.tpl", nil)
		panic(err)
	}
}

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
}

func index(c *gin.Context) {
	rows, err := GetTablesStmt.QueryContext(c)
	checkErr(c, err)
	defer rows.Close()

	var tables []Table
	for rows.Next() {
		var table Table
		err = rows.Scan(&table.Name, &table.Comment)
		checkErr(c, err)
		tables = append(tables, table)
	}

	c.HTML(200, "index.tpl", tables)
}

func getTable(c *gin.Context) {
	tableName := c.Query("tableName")
	rows, err := gDb.QueryContext(c, GetTableSql, config.DbCfg().LibName, tableName)
	checkErr(c, err)
	defer rows.Close()

	var columns []Column
	for rows.Next() {
		var column Column
		err = rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnType, &column.ColumnComment)
		checkErr(c, err)
		columns = append(columns, column)
	}
	table := Table{
		Name:    tableName,
		Columns: columns,
	}

	c.HTML(200, "getTable.tpl", table)
}

type TableRecord struct {
	Columns []string
	Records []Record
}

type Record struct {
	Values []string
}

func getTableRecords(c *gin.Context) {
	tableName := c.Query("tableName")
	rows, err := gDb.QueryContext(c, GetTableSql, config.DbCfg().LibName, tableName)
	checkErr(c, err)
	defer rows.Close()

	var columns []Column
	for rows.Next() {
		var column Column
		err = rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnType, &column.ColumnComment)
		checkErr(c, err)
		columns = append(columns, column)
	}

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
	checkErr(c, err)
	defer rowRecords.Close()

	for rowRecords.Next() {

		err = rowRecords.Scan(columnVals...)
		checkErr(c, err)
		var record Record
		for _, realColumnVal := range realColumnVals {
			record.Values = append(record.Values, string(realColumnVal))
		}
		tableRecord.Records = append(tableRecord.Records, record)
	}

	c.HTML(200, "getTableRecords.tpl", tableRecord)
}

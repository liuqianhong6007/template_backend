package route

import (
	"github.com/liuqianhong6007/template_backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
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
	tableName := c.Param("tableName")
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

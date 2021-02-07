package route

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/template_backend/config"
	"github.com/liuqianhong6007/template_backend/db"
)

func init() {
	AddRoute(Routes{
		{
			Method:  http.MethodGet,
			Path:    "/index",
			Handler: Index,
		},
		{
			Method:  http.MethodGet,
			Path:    "/getTables",
			Handler: GetTables,
		},
		{
			Method:  http.MethodGet,
			Path:    "/getTable",
			Handler: GetTable,
		},
		{
			Method:  http.MethodGet,
			Path:    "/getTableRecords",
			Handler: GetTableRecords,
		},
		{
			Method:  http.MethodPost,
			Path:    "/updateTableRecord",
			Handler: UpdateTableRecord,
		},
		{
			Method:  http.MethodPost,
			Path:    "/deleteTableRecord",
			Handler: DeleteTableRecord,
		},
	})
}

type Routes []Route

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var routeMap = make(map[string]Route)

func AddRoute(routes Routes) {
	for _, route := range routes {
		if _, ok := routeMap[route.Path]; ok {
			panic("duplicate register router: " + route.Path)
		}
		routeMap[route.Path] = route
	}
}

func RegisterRoute(engine *gin.Engine) {
	// load template file
	engine.LoadHTMLGlob("tpl/**/*")

	for _, route := range routeMap {
		engine.Handle(route.Method, route.Path, route.Handler)
	}
	// 404 page
	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tpl", nil)
	})

	// static file server
	engine.StaticFS("/static", http.Dir("./static"))
}

var gDb *sql.DB

func InitDatabase() {
	dbCfg := config.DbCfg()
	gDb = db.InitDatabase(db.Config{
		Driver:  "mysql",
		Addr:    dbCfg.Host,
		Port:    dbCfg.Port,
		User:    dbCfg.User,
		Pass:    dbCfg.Pass,
		LibName: dbCfg.LibName,
	})

	prepareStmt(gDb)
}

var (
	GetTablesStmt *sql.Stmt
	GetTableSql   = `select COLUMN_NAME,DATA_TYPE,COLUMN_TYPE,COLUMN_COMMENT,COLUMN_KEY from information_schema.COLUMNS where TABLE_SCHEMA = ? AND TABLE_NAME = ?`
)

func prepareStmt(db *sql.DB) {
	var err error
	GetTablesStmt, err = db.Prepare(fmt.Sprintf(`select table_name,table_comment from information_schema.TABLES where TABLE_SCHEMA = '%s'`, config.DbCfg().LibName))
	if err != nil {
		panic(err)
	}
}

func BuildSql(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func checkValue(c *gin.Context, checkValue interface{}, errMsg ...string) {
	switch val := checkValue.(type) {
	case error:
		if val != nil {
			c.HTML(http.StatusInternalServerError, "500.tpl", nil)
			panic(strings.Join(errMsg, "\n") + "\n" + val.Error())
		}
	case bool:
		if !val {
			c.HTML(http.StatusInternalServerError, "500.tpl", nil)
			panic(strings.Join(errMsg, "\n"))
		}
	}
}

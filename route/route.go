package route

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/template_backend/config"
	"github.com/liuqianhong6007/template_backend/db"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var routes = make(map[string]Route)

func AddRoute(route Route) {
	if _, ok := routes[route.Path]; ok {
		panic("duplicate register router: " + route.Path)
	}
	routes[route.Path] = route
}

func RegisterRoute(engine *gin.Engine) {
	for _, route := range routes {
		engine.Handle(route.Method, route.Path, route.Handler)
	}
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
	GetTableSql   = `select COLUMN_NAME,DATA_TYPE,COLUMN_TYPE,COLUMN_COMMENT from information_schema.COLUMNS where TABLE_SCHEMA = ? AND TABLE_NAME = ?`
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

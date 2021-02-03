package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/template_backend/config"
	"github.com/liuqianhong6007/template_backend/db"
	"github.com/liuqianhong6007/template_backend/route"
)

func main() {
	initContext()

	r := gin.Default()
	r.LoadHTMLGlob("tpl/*")
	route.RegisterRoute(r)

	serverAddr := fmt.Sprintf("%s:%d", config.ServerCfg().Host, config.ServerCfg().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}

func initContext() {
	dbCfg := config.DbCfg()
	db.InitDatabase(db.Config{
		Driver:  "mysql",
		Addr:    dbCfg.Host,
		Port:    dbCfg.Port,
		User:    dbCfg.User,
		Pass:    dbCfg.Pass,
		LibName: dbCfg.LibName,
	})
}

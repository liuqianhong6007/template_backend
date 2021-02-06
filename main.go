package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/template_backend/config"
	"github.com/liuqianhong6007/template_backend/route"
)

func main() {
	route.InitDatabase()

	r := gin.Default()
	route.RegisterRoute(r)

	serverAddr := fmt.Sprintf("%s:%d", config.ServerCfg().Host, config.ServerCfg().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}

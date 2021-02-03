package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liuqianhong6007/template_backend/route"

	_ "github.com/liuqianhong6007/template_backend/route"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tpl/*")
	route.RegisterRoute(r)
	r.Run(":8081")
}

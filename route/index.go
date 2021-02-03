package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	AddRoute(Route{
		Method:  http.MethodGet,
		Path:    "/index",
		Handler: index,
	})
}

func index(c *gin.Context) {
	c.HTML(200, "index.tpl", "")
}

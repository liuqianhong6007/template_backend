package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	AddRoute(Route{
		Method:  http.MethodGet,
		Path:    "/ping",
		Handler: ping,
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

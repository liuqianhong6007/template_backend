package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	AddRoute(Routes{
		{
			Method:  http.MethodGet,
			Path:    "/ping",
			Handler: ping,
		},
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

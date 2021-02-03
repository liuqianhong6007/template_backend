package route

import "github.com/gin-gonic/gin"

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

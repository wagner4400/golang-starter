// internal/entrypoint/http/router/router.go
package router

import (
	"github.com/gin-gonic/gin"
)

type RouteHandler interface {
	RegisterRoutes(*gin.RouterGroup)
}

type Router struct {
	engine     *gin.Engine
	middleware []gin.HandlerFunc
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine:     engine,
		middleware: make([]gin.HandlerFunc, 0),
	}
}

func (r *Router) AddMiddleware(middleware ...gin.HandlerFunc) {
	r.middleware = append(r.middleware, middleware...)
}

func (r *Router) ConfigureRoutes(handlers ...RouteHandler) {

	// Apply global middleware
	for _, m := range r.middleware {
		r.engine.Use(m)
	}

	// Create base API group
	api := r.engine.Group("/api/v1")

	// Register routes for each handler
	for _, handler := range handlers {
		handler.RegisterRoutes(api)
	}
}

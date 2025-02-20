package author

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/infrastructure"
)

type Route struct {
	router      *infrastructure.Router
	controller  *Controller
	groupRouter *gin.RouterGroup
}

func NewRoute(router *infrastructure.Router, controller *Controller) *Route {
	route := Route{router: router, controller: controller}
	route.groupRouter = route.router.Group("api/authors")
	return &route
}

func RegisterRoute(r *Route) {
	r.groupRouter.GET("", r.controller.HandleFindAuthors)
	r.groupRouter.POST("", r.controller.HandleAddAuthor)
	r.groupRouter.GET("/filter", r.controller.HandleFilterAuthors)
	r.groupRouter.GET("/:id", r.controller.HandleGetAuthorByID)
	r.groupRouter.POST("/books", r.controller.HandleAddAuthorWithBooks)
}

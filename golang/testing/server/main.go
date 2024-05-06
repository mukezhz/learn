package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func SetupRouter(engine *gin.Engine) {
	engine.GET("/ping", PingHandler)
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func HTTPHandler(handler gin.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, _ := gin.CreateTestContext(w)
		handler(c)
	}
}

type RouterParams struct {
	fx.In
	Engine *gin.Engine
}

func NewApp(routerParams RouterParams) {
	// Application setup code using router
	fmt.Println("NEW::APP::", routerParams)
	fmt.Println("LENGTH::", len(routerParams.Engine.Routes()))
	routerParams.Engine.Run(":8888")
}

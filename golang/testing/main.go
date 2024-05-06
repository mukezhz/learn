package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/learn/golang/testing/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(func() *gin.Engine {
			engine := gin.Default()
			server.SetupRouter(engine)
			return engine
		}),
		fx.Invoke(server.NewApp),
	)

	app.Run()
}

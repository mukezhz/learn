package main_test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mukezhz/learn/golang/testing/server"
	"github.com/onsi/ginkgo/v2"
	"github.com/steinfletcher/apitest"
	"go.uber.org/fx"
)

var _ = ginkgo.Describe("PingHandler", func() {
	var (
		router *gin.Engine
	)

	ginkgo.BeforeEach(func() {
		fmt.Println("BeforeEach")
		router = gin.Default()
		server.SetupRouter(router)
	})

	ginkgo.BeforeEach(func(ctx ginkgo.SpecContext) {
		app := fx.New(
			fx.Provide(func() *gin.Engine {
				engine := gin.Default()
				server.SetupRouter(engine)
				return engine
			}),
			fx.Invoke(func(engine *gin.Engine) {
				go func() {
					defer ginkgo.GinkgoRecover()
				}()
			}),
		)

		go app.Run()
	})

	ginkgo.It("should return pong for GET /ping (apitest)", func() {
		handler := server.PingHandler

		apitest.New().
			Handler(server.HTTPHandler(handler)).
			Get("/ping").
			Expect(ginkgo.GinkgoT()).
			Status(http.StatusOK).
			Body(`{"message":"pong"}`).
			End()
	})
})

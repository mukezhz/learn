package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from AWS Lambda with Gin!"})
	})

	type paylaod struct {
		Name string `json:"name"`
	}
	r.POST("/ping", func(c *gin.Context) {
		log.Println("hello")
		r := paylaod{}
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "hello: " + r.Name})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	ginLambda = ginadapter.NewV2(r)
}

func main() {
	lambda.Start(ginLambda.ProxyWithContext)
}

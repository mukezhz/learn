package main

import (
	"log"
	"mukezhz/openapi-example/server_gen"
	"mukezhz/openapi-example/server_gen/user_gen"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Register handlers
	userHandler := server_gen.NewUserHandler()
	// orgHandler := server_gen.NewOrganizationHandler()

	user_gen.RegisterHandlers(router, userHandler)
	// orgs.RegisterHandlers(router, orgHandler)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorType struct {
	Field      string
	Message    string
	StatusCode int
}

type CustomError struct {
	Errors []ErrorType
}

func (e ErrorType) Error() string {
	return e.Message
}

func (e CustomError) Error() string {
	var messages []string
	for _, err := range e.Errors {
		messages = append(messages, err.Message)
	}
	return fmt.Sprintf("Custom errors: %v", messages)
}

func CustomErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := c.Errors.Last()
			if err != nil {
				statusCode := http.StatusInternalServerError
				var errors []gin.H

				switch e := err.Err.(type) {
				case CustomError:
					statusCode = getStatusCode(e)
					errors = getErrors(e)

				default:
					errors = append(errors, gin.H{"error": "Internal Server Error"})
					fmt.Println("Error:", err.Error())
				}

				c.AbortWithStatusJSON(statusCode, gin.H{"errors": errors})
				fmt.Printf("Custom error with status code %d: %v\n", statusCode, errors)
			}
		}()

		c.Next()
	}
}

func getStatusCode(e CustomError) int {
	if len(e.Errors) > 0 {
		return e.Errors[0].StatusCode
	}
	return http.StatusInternalServerError
}

func getErrors(e CustomError) []gin.H {
	var errors []gin.H
	for _, err := range e.Errors {
		errorObj := gin.H{"message": err.Message}
		if err.Field != "" {
			errorObj["field"] = err.Field
		}
		errors = append(errors, errorObj)
	}
	return errors
}

func main() {
	router := gin.Default()

	router.Use(CustomErrorHandler())

	router.GET("/ping", func(c *gin.Context) {
		c.Error(CustomError{Errors: []ErrorType{
			{Message: "Custom error", Field: "email", StatusCode: http.StatusBadRequest},
			{Message: "Custom error", Field: "", StatusCode: http.StatusBadRequest},
		}})
	})

	router.Run(":8080")
}

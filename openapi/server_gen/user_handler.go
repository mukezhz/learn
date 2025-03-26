package server_gen

import (
	"mukezhz/openapi-example/server_gen/user_gen"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler implements the generated UserServerInterface
type UserHandler struct {
	users []user_gen.User
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		users: []user_gen.User{},
	}
}

// GetUsers returns a list of users
func (h *UserHandler) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.users)
}

// CreateUser adds a new user
func (h *UserHandler) CreateUser(ctx *gin.Context) {
	b := user_gen.User{}
	ctx.ShouldBindJSON(&b)
	h.users = append(h.users, b)
	ctx.JSON(http.StatusCreated, b)
}

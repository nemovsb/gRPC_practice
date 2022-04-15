package controllers

import (
	"github.com/gin-gonic/gin"
)

type HandlerSet struct {
	UsersHandler
}

type UsersHandler interface {

	// Get user by id
	Get(ctx *gin.Context)

	// Update user by id
	Put(ctx *gin.Context)

	// Create user
	Post(ctx *gin.Context)

	// Delete user by id
	Del(ctx *gin.Context)
}

func NewHandlerSet(user UsersHandler) *HandlerSet {
	return &HandlerSet{
		UsersHandler: user,
	}
}

func NewRouter(h HandlerSet) (router *gin.Engine) {
	router = gin.Default()

	user := router.Group("/users")
	{
		user.GET("/{id}", h.UsersHandler.Get)
		user.PUT("/{id}", h.UsersHandler.Put)
		user.POST("", h.UsersHandler.Post)
		user.DELETE("/{id}", h.UsersHandler.Del)
	}

	return router
}

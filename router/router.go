package router

import (
	"blogging-platform-api/handler"
	"blogging-platform-api/initialize"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, app *initialize.Application) {
	handler := handler.NewHandler(app)
	router.GET("/models", handler.Get)
	router.POST("/post", handler.CreatePost)
	router.GET("/generation/:id", handler.Get)
}

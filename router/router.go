package router

import (
	"blogging-platform-api/handler"
	"blogging-platform-api/initialize"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, app *initialize.Application) {
	handler := handler.NewHandler(app)

	postsGroup := router.Group("/posts")
	{
		postsGroup.GET("/:id", handler.GetDetailPost)
		postsGroup.POST("", handler.CreatePost)
	}
}

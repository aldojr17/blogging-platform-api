package middleware

import (
	"blogging-platform-api/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NoRouteMiddleware(ctx *gin.Context) {
	handler.ResponseNotFound(ctx, fmt.Errorf("not found"))
}

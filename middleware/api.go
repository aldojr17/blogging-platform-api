package middleware

import (
	log "blogging-platform-api/logger"
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func APIMiddleware(ctx *gin.Context) {
	start := time.Now()

	// Log request details
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Infof("Error reading request body: %v", err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // Restore request body

	log.Info("Request:", "\n Method", ctx.Request.Method, "\n URL", ctx.Request.URL, "\n Body", string(body))

	// Process request
	ctx.Next()

	// Log response details
	duration := time.Since(start)
	statusCode := ctx.Writer.Status()

	log.Info("Response:", "\nStatus", statusCode, "\nDuration", duration)
}

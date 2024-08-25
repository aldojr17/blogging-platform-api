package main

import (
	"blogging-platform-api/initialize"
	"blogging-platform-api/middleware"
	"blogging-platform-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := initialize.InitApp()

	r := gin.New()
	r.NoRoute(middleware.NoRouteMiddleware)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.APIMiddleware)

	router.Routes(r, app)

	if err := r.Run(); err != nil {
		panic(err)
	}
}

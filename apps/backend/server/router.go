package main

import (
	"ai-saas-schedular-server/server/api"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func InitializeRoutes(router *gin.Engine) {
	api.SetupAPIRoutes(router.Group("/api"))
}

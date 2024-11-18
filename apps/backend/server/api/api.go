package api

import (
	v1 "ai-saas-schedular-server/server/api/v1"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.RouterGroup) {
	v1.SetupV1Routes(router.Group("/v1"))
}

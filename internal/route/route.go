package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/middleware"
)

func Init() *gin.Engine {
	router := gin.New()

	router.Use(gin.CustomRecovery(middleware.ExceptionHandler))

	v1 := router.Group("/v1")

	handleUser(v1)

	return router
}

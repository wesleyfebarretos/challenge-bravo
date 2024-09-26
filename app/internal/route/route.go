package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/middleware"
)

func Init() *gin.Engine {
	router := gin.New()
	v1 := router.Group("/v1")

	handleSwagger(v1)

	router.Use(middleware.Log)
	router.Use(gin.CustomRecovery(middleware.ExceptionHandler))

	handleUser(v1)
	handleCurrency(v1)
	handleAuth(v1)

	return router
}

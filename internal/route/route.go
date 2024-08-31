package route

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/v1")

	handleUser(v1)

	return router
}

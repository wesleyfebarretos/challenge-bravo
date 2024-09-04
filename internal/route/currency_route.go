package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/middleware"
)

func handleCurrency(router *gin.RouterGroup) {
	currencyRoute := router.Group("currency")

	currencyRoute.Use(middleware.Jwt)

	currencyRoute.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})
}

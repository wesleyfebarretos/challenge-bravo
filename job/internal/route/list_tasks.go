package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/middleware"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/scheduler"
)

func handleListTasks(router *gin.RouterGroup) {
	listTaskRoute := router.Group("")

	listTaskRoute.Use(middleware.Jwt)

	listTaskRoute.GET("tasks", func(c *gin.Context) {
		scheduler := scheduler.New()

		c.JSON(http.StatusOK, scheduler.GetAllTasks())
	})
}

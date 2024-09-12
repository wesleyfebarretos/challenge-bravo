package route

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/middleware"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/scheduler"
)

func handleRemoveTask(route *gin.RouterGroup) {
	removeTaskRoute := route.Group("")

	removeTaskRoute.Use(middleware.Jwt)

	removeTaskRoute.DELETE("tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusBadRequest,
				"message": "id param is required and need to be a positive and integer number",
			})
			return
		}

		scheduler.New().RemoveTask(cron.EntryID(id))

		c.JSON(http.StatusOK, true)
	})
}

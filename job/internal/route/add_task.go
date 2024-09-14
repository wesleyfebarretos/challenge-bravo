package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/enum"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/middleware"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/task"
)

type AddTaskRequest struct {
	Name string `json:"name"`
}

func handleAddTask(router *gin.RouterGroup) {
	listTaskRoute := router.Group("")

	listTaskRoute.Use(middleware.Jwt)

	listTaskRoute.POST("tasks", func(c *gin.Context) {
		body := AddTaskRequest{}

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		if body.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "task name is required",
			})
			return

		}

		switch body.Name {
		case enum.CurrencyUpdaterTask:
			task.NewCurrencyUpdater().AddToScheduler()
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "task not available",
			})
			return
		}

		c.JSON(http.StatusOK, true)
	})
}

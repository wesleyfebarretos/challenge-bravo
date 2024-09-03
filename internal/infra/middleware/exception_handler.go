package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

func ExceptionHandler(c *gin.Context, recovered any) {
	if exception, ok := recovered.(*exception.HttpException); ok {
		c.JSON(exception.Code, gin.H{"code": exception.Code, "message": exception.Message})
	} else {
		log.Printf("Exception not mapped: %s", recovered)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}


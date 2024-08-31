package handler

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

func getIdFromReq(c *gin.Context) int32 {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		panic(exception.BadRequest(fmt.Sprintf("invalid id parameter %s", id)))
	}

	return int32(intId)
}

func getUuidFromReq(c *gin.Context) uuid.UUID {
	uuidRequest := c.Param("uuid")

	parsedUuid, err := uuid.Parse(uuidRequest)
	if err != nil {
		panic(exception.BadRequest(fmt.Sprintf("invalid uuid parameter %s", uuidRequest)))
	}

	return parsedUuid
}

func readBody[B any](c *gin.Context, body *B) {
	err := c.ShouldBindJSON(body)
	if err == io.EOF {
		panic(exception.BadRequest("empty request body"))
	}
	if err != nil {
		panic(exception.BadRequest(err.Error()))
	}
}

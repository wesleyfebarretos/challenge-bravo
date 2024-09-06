package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/internal/usecase"
)

type DeleteCurrencyHandler struct {
	useCase usecase.DeleteCurrencyUseCase
}

func (h DeleteCurrencyHandler) Execute(c *gin.Context) {
	id := getIdFromReq(c)

	h.useCase.Execute(c, id)

	c.JSON(http.StatusOK, true)
}

func NewDeleteCurrencyHandler(useCase usecase.DeleteCurrencyUseCase) DeleteCurrencyHandler {
	return DeleteCurrencyHandler{
		useCase: useCase,
	}
}

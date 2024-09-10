package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/exception"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/usecase"
)

type CurrencyConversionHandler struct {
	useCase usecase.CurrencyConversionUseCase
}

type CurrencyConversionRequest struct {
	From   string  `form:"from" example:"usd"`
	To     string  `form:"to" example:"brl"`
	Amount float64 `form:"amount" example:"10.2"`
}

func (c CurrencyConversionRequest) Valid() error {
	reqError := []string{}

	if c.From == "" {
		reqError = append(reqError, "from is required")
	}

	if c.To == "" {
		reqError = append(reqError, "to is required")
	}

	if c.Amount == 0 {
		reqError = append(reqError, "amount is required")
	}

	if len(reqError) > 0 {
		return fmt.Errorf("query param(s) %s", strings.Join(reqError, ", "))
	}

	return nil
}

type CurrencyConversionResponse struct {
	Label string  `json:"label" example:"5.57 BRL"`
	Value float64 `json:"value" example:"5.57"`
}

func (c CurrencyConversionResponse) MapToResponse(p usecase.CurrencyConversionDTO) CurrencyConversionResponse {
	return CurrencyConversionResponse{
		Label: p.Label,
		Value: p.Value,
	}
}

func (h CurrencyConversionHandler) Execute(c *gin.Context) {
	queryParams := CurrencyConversionRequest{}

	c.ShouldBindQuery(&queryParams)

	err := queryParams.Valid()
	if err != nil {
		panic(exception.BadRequest(err.Error()))
	}

	conversion := h.useCase.Execute(c, queryParams.From, queryParams.To, queryParams.Amount)

	res := CurrencyConversionResponse{}

	c.JSON(http.StatusCreated, res.MapToResponse(conversion))
}

func NewCurrencyConversionHandler(useCase usecase.CurrencyConversionUseCase) CurrencyConversionHandler {
	return CurrencyConversionHandler{
		useCase: useCase,
	}
}

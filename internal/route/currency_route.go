package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/middleware"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/repository/currency_repository"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/web/handler"
	"github.com/wesleyfebarretos/challenge-bravo/internal/usecase"
)

func handleCurrency(router *gin.RouterGroup) {
	currencyRoute := router.Group("currency")

	currencyRoute.Use(middleware.Jwt)

	createCurrencyHandler := handler.NewCreateCurrencyHandler(usecase.NewCreateCurrencyUseCase(currency_repository.New()))

	currencyRoute.POST("", createCurrencyHandler.Execute)
}

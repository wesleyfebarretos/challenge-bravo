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

	repository := currency_repository.New()

	createCurrencyHandler := handler.NewCreateCurrencyHandler(usecase.NewCreateCurrencyUseCase(repository))
	updateCurrencyHandler := handler.NewUpdateCurrencyHandler(usecase.NewUpdateCurrencyUseCase(repository))
	deleteCurrencyHandler := handler.NewDeleteCurrencyHandler(usecase.NewDeleteCurrencyUseCase(repository))
	FindCurrencyByCodeHandler := handler.NewFindCurrencyByCodeHandler(usecase.NewFindCurrencyByCodeUseCase(repository))
	FindCurrencyByIdHandler := handler.NewFindCurrencyByIdHandler(usecase.NewFindCurrencyByIdUseCase(repository))
	FindAllCurrencyHandler := handler.NewFindAllCurrencyHandler(usecase.NewFindAllCurrencyUseCase(repository))
	currencyConversionHandler := handler.NewCurrencyConversionHandler(usecase.NewCurrencyConversionUseCase())

	currencyRoute.POST("", createCurrencyHandler.Execute)
	currencyRoute.PUT(":id", updateCurrencyHandler.Execute)
	currencyRoute.DELETE(":id", deleteCurrencyHandler.Execute)
	currencyRoute.GET("", FindAllCurrencyHandler.Execute)
	currencyRoute.GET(":id", FindCurrencyByIdHandler.Execute)
	currencyRoute.GET("/code/:code", FindCurrencyByCodeHandler.Execute)
	currencyRoute.GET("convert", currencyConversionHandler.Execute)
}

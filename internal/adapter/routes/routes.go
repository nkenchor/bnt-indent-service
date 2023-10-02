package routes

import (
	"bnt/bnt-indent-service/internal/adapter/api"
	errorhelper "bnt/bnt-indent-service/internal/core/helper/error-helper"
	logger "bnt/bnt-indent-service/internal/core/helper/log-helper"
	message "bnt/bnt-indent-service/internal/core/helper/message-helper"
	"bnt/bnt-indent-service/internal/core/middleware"
	"bnt/bnt-indent-service/internal/core/services"
	ports "bnt/bnt-indent-service/internal/port"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	indentRepository ports.IndentRepository) *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	indentService := services.NewIndent(indentRepository)

	handler := api.NewHTTPHandler(
		indentService)

	logger.LogEvent("INFO", "Configuring Routes!")
	router.Use(middleware.LogRequest)

	//router.Use(middleware.SetHeaders)

	router.Group("/api/indent")
	{
		router.POST("/api/indent", handler.CreateIndent)
		router.PUT("/api/indent/:reference", handler.UpdateIndent)
		router.PUT("/api/indent/disable/:reference", handler.DisableIndent)
		router.PUT("/api/indent/enable/:reference", handler.EnableIndent)
		router.DELETE("/api/indent/:reference", handler.DeleteIndent)
		router.GET("/api/indent/page/:page", handler.GetIndentList)
		router.GET("/api/indent/:reference", handler.GetIndentByRef)

	}
	router.Group("/api/platform")
	{
		router.GET("/api/platform/year/:count", handler.GetYearList)
		router.GET("/api/platform/denomination/list", handler.GetDenominationList)
		router.GET("/api/platform/location/list", handler.GetLocationList)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			errorhelper.ErrorMessage(errorhelper.NoResourceError, message.NoResourceFound))
	})

	return router
}

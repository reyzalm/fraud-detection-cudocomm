package http

import (
	"github.com/CudoCommunication/cudocomm/internal/middleware"
	"github.com/labstack/echo/v4"
)

func MapFraudRoutes(apiGroup *echo.Group, h FraudHandlers, mw *middleware.MiddlewareManager) {

	fraudGroup := apiGroup.Group("/fraud-detection")
	fraudGroup.GET("/users/:user_id/transactions/:transaction_id", h.DetectFraud(), mw.AuthJWTMiddleware())
}

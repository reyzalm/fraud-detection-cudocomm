package auth

import (
	"github.com/CudoCommunication/cudocomm/internal/middleware"
	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(authGroup *echo.Group, h AuthHandlers, mw *middleware.MiddlewareManager) {
	authGroup.POST("/login", h.Login())
}

package server

import (
	"net/http"

	gorm_repository "github.com/CudoCommunication/cudocomm/internal/database/gorm"
	"github.com/CudoCommunication/cudocomm/internal/logger"
	apiMiddleware "github.com/CudoCommunication/cudocomm/internal/middleware"
	authHttp "github.com/CudoCommunication/cudocomm/internal/modules/auth/delivery/http"
	authUseCase "github.com/CudoCommunication/cudocomm/internal/modules/auth/usecase"
	fraudHttp "github.com/CudoCommunication/cudocomm/internal/modules/fraud/delivery/http"
	fraudUseCase "github.com/CudoCommunication/cudocomm/internal/modules/fraud/usecase"
	"github.com/CudoCommunication/cudocomm/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	userRepo := gorm_repository.NewUserRepositoryGorm(s.db)
	transRepo := gorm_repository.NewTransactionRepositoryGorm(s.db)

	logger := logger.NewEchoLogger(s.echo.Logger)
	authUC := authUseCase.NewAuthUseCase(logger, userRepo)
	fraudUC := fraudUseCase.NewFraudUseCase(logger, transRepo)

	authHandler := authHttp.NewAuthHandler(authUC)
	fraudHandler := fraudHttp.NewFraudHandler(fraudUC)

	mw := apiMiddleware.NewMiddlewareManager()

	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("20M"))
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandler, mw)
	fraudHttp.MapFraudRoutes(v1, fraudHandler, mw)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, utils.Response{
			Success: true,
			Message: "OK!",
		})
	})

	return nil
}

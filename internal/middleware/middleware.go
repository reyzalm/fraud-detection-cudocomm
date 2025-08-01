package middleware

import (
	"net/http"

	"github.com/CudoCommunication/cudocomm/config"
	"github.com/CudoCommunication/cudocomm/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type MiddlewareManager struct {
}

func NewMiddlewareManager() *MiddlewareManager {
	return &MiddlewareManager{}
}

func (m *MiddlewareManager) AuthJWTMiddleware() echo.MiddlewareFunc {
	jwtConfig := echojwt.Config{

		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utils.JwtCustomClaims)
		},

		SigningKey: []byte(config.Env.SecretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Unauthorized: " + err.Error(),
			})
		},
		SuccessHandler: func(c echo.Context) {
			if user, ok := c.Get("user").(jwt.MapClaims); ok {
				
				if idStr, idOk := user["id"].(string); idOk {
					c.Set("id", idStr)
				}
			}
		},
	}

	return echojwt.WithConfig(jwtConfig)
}

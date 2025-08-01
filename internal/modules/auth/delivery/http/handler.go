package auth

import (
	"net/http"

	"github.com/CudoCommunication/cudocomm/config"
	auth "github.com/CudoCommunication/cudocomm/internal/modules/auth/usecase"
	"github.com/CudoCommunication/cudocomm/pkg/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandlers interface {
	Login() echo.HandlerFunc
}

type authHandlersImpl struct {
	authUC auth.AuthUseCase
}

func NewAuthHandler(authUC auth.AuthUseCase) AuthHandlers {
	return &authHandlersImpl{
		authUC: authUC,
	}
}

func (h *authHandlersImpl) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(auth.LoginDTO)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, utils.Response{Message: config.ERROR_BINDING_REQUESST})
		}

		if err := utils.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
		}

		data, err := h.authUC.Login(*req)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.Response{Message: err.Error()})
		}

		return c.JSON(http.StatusOK, utils.Response{
			Success: true,
			Message: "Login successful",
			Data:    data,
		})
	}
}

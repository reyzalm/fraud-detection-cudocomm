package http

import (
	"net/http"
	"strconv"

	"github.com/CudoCommunication/cudocomm/internal/modules/fraud/usecase"
	"github.com/CudoCommunication/cudocomm/pkg/utils"
	"github.com/labstack/echo/v4"
)

type FraudHandlers interface {
	DetectFraud() echo.HandlerFunc
}

type fraudHandlersImpl struct {
	fraudUC usecase.FraudUseCase
}

func NewFraudHandler(fraudUC usecase.FraudUseCase) FraudHandlers {
	return &fraudHandlersImpl{fraudUC: fraudUC}
}

func (h *fraudHandlersImpl) DetectFraud() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.Response{Message: "Invalid User ID"})
		}

		transactionID, err := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.Response{Message: "Invalid Transaction ID"})
		}

		result, err := h.fraudUC.DetectFraud(userID, transactionID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to process fraud detection: " + err.Error()})
		}

		return c.JSON(http.StatusOK, utils.Response{
			Success: true,
			Message: "Fraud detection process completed",
			Data:    result,
		})
	}
}

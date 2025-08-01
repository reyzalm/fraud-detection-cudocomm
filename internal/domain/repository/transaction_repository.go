package repository

import (
	"time"

	"github.com/CudoCommunication/cudocomm/internal/domain/models"
)

type TransactionRepository interface {
	GetTransactionByID(transactionID int64) (*models.Transaction, error)
	GetUserTransactions(userID int64, untilDate time.Time) ([]models.Transaction, error)
}

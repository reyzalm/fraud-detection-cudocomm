package gorm_repository

import (
	"time"

	"github.com/CudoCommunication/cudocomm/internal/domain/models"
	"github.com/CudoCommunication/cudocomm/internal/domain/repository"
	"gorm.io/gorm"
)

type transactionRepositoryGorm struct {
	db *gorm.DB
}

func NewTransactionRepositoryGorm(db *gorm.DB) repository.TransactionRepository {
	return &transactionRepositoryGorm{
		db: db,
	}
}

func (r *transactionRepositoryGorm) GetTransactionByID(transactionID int64) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.First(&transaction, transactionID).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepositoryGorm) GetUserTransactions(userID int64, untilDate time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id = ? AND transaction_date < ?", userID, untilDate).
		Order("transaction_date desc").
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}
	return transactions, nil
}

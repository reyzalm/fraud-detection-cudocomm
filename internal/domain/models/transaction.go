package models

import (
	"time"
)

type Transaction struct {
	ID              int64     `json:"id" gorm:"primaryKey"`
	UserID          int64     `json:"user_id"`
	OrderID         string    `json:"order_id"`
	TransactionDate time.Time `json:"transaction_date"`
	Amount          float64   `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}

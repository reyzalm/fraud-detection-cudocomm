package repository

import (
	"github.com/CudoCommunication/cudocomm/internal/domain/models"
)

type UserRepository interface {
	GetByField(field string, value any) (*models.User, error)
}

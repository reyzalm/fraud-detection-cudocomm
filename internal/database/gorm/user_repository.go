package gorm_repository

import (
	"fmt"
	"github.com/CudoCommunication/cudocomm/internal/domain/models"
	"github.com/CudoCommunication/cudocomm/internal/domain/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryGorm{
		db: db,
	}
}

func (ur *UserRepositoryGorm) GetByField(field string, value any) (*models.User, error) {
	var user models.User
	if err := ur.db.Where(fmt.Sprintf("%s = ?", field), value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepositoryGorm) GetById(id *uuid.UUID) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

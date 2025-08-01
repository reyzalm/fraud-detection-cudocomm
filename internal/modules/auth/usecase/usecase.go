package auth

import (
	"errors"

	"github.com/CudoCommunication/cudocomm/config"
	"github.com/CudoCommunication/cudocomm/internal/domain"
	"github.com/CudoCommunication/cudocomm/internal/domain/models"
	"github.com/CudoCommunication/cudocomm/internal/domain/repository"
	"github.com/CudoCommunication/cudocomm/pkg/utils"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	Login(req LoginDTO) (*models.UserWithToken, error)
}

type authUseCaseImpl struct {
	logger   domain.Logger
	userRepo repository.UserRepository
}

func NewAuthUseCase(logger domain.Logger, userRepo repository.UserRepository) AuthUseCase {
	return &authUseCaseImpl{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (u *authUseCaseImpl) Login(req LoginDTO) (*models.UserWithToken, error) {
	
	user, err := u.userRepo.GetByField("email", *req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(config.ERROR_LOGIN)
		}
		u.logger.Error(&domain.LoggerPayload{Loc: "auth.Login.GetByField", Msg: err.Error()})
		return nil, errors.New(config.ERROR_DATABASE)
	}

	if err := user.ComparePasswords(req.Password); err != nil {
		return nil, errors.New(config.ERROR_LOGIN)
	}

	token, err := utils.GenerateJWTToken(user)
	if err != nil {
		u.logger.Error(&domain.LoggerPayload{Loc: "auth.Login.GenerateToken", Msg: err.Error()})
		return nil, errors.New("error generating token")
	}

	return &models.UserWithToken{
		User:  user,
		Token: token,
	}, nil
}

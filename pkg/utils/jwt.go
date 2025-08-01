package utils

import (
	"time"

	"github.com/CudoCommunication/cudocomm/config"
	"github.com/CudoCommunication/cudocomm/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(user *models.User) (string, error) {

	claims := &JwtCustomClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.Env.JwtExpired))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.Env.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

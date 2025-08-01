package models

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int64      `json:"id" gorm:"primaryKey"`
	Name            *string    `json:"name"`
	Email           *string    `json:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        *string    `json:"-"`
	RememberToken   *string    `json:"remember_token,omitempty"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

func (u *User) HashPassword() error {
	if u.Password == nil {
		return errors.New("password is not provided")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password *string) error {
	if u.Password == nil || password == nil {
		return errors.New("invalid credentials")
	}

	hashFromDB := *u.Password
	if strings.HasPrefix(hashFromDB, "$2y$") {
		hashFromDB = "$2a$" + strings.TrimPrefix(hashFromDB, "$2y$")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashFromDB), []byte(*password))
	return err
}

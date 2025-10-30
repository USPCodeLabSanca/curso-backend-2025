package user

import (
	"codelab/backend/internal/domain"
	"codelab/backend/pkg/token"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
Verifica se os dados informados estão corretos
e batem com os armazenados na base.

Em caso de sucesso, um token de acesso de 24h
é gerado e retornado.
*/

type LoginUserUseCase struct {
	db    *gorm.DB
	token token.TokenDoer
}

func NewLoginUserUseCase(db *gorm.DB, token token.TokenDoer) *LoginUserUseCase {
	return &LoginUserUseCase{db: db, token: token}
}

func (usecase *LoginUserUseCase) Execute(email, password string) (string, error) {
	ctx := context.Background()

	user, err := gorm.G[domain.User](usecase.db).Where("email = ?", email).First(ctx)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Verificar senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid user or password")
	}

	// Gerar token com expiração de 24h
	tokenData := map[string]any{
		"id":    user.ID,
	}

	tokenStr, err := usecase.token.Encrypt(tokenData, time.Now().Add(24*time.Hour))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

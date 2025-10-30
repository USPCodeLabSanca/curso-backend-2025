package user

import (
	"codelab/backend/internal/domain"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
Caos de uso: criar usuário.

O email deve ser verificado para garantir que
seja único.

O usuário criado é salvo na base.
A senha DEVE ser criptografada.
*/

type CreateUserUseCase struct {
	db *gorm.DB
}

func NewCreateUserUseCase(db *gorm.DB) *CreateUserUseCase {
	return &CreateUserUseCase{db}
}

func (usecase *CreateUserUseCase) Execute(payload *domain.CreateUserDTO) error {
	ctx := context.Background()

	// Verifica se o email já está cadastrado
	_, err := gorm.G[domain.User](usecase.db).Where("email = ?", payload.Email).First(ctx)
	if err == nil {
		return errors.New("email already in use")
	}

	// Criptografar a senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Criar o usuário
	user := domain.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
	}

	return gorm.G[domain.User](usecase.db).Create(ctx, &user)
}
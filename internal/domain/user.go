package domain

import "time"

/*
Entidade "Usuário", definida usando o modelo
padrão do GORM e os atributos.

Note que o email deve ser único.
*/
type User struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at,omitempty"`
	Name      string        `json:"name"`
	Email     string        `json:"email" gorm:"uniqueIndex"`
	Password  string        `json:"password"`
}

type CreateUserDTO struct {
	Name      string        `json:"name" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Password  string        `json:"password" binding:"required"`
}

type LoginUserDTO struct {
	Email     string        `json:"email" binding:"required"`
	Password  string        `json:"password" binding:"required"`
}
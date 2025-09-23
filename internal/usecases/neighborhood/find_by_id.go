package neighborhood

import (
	"codelab/backend/internal/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

/*
Caso de Uso:
Buscar um bairro por ID.

O ID é informado e, caso o bairro exista,
os dados são retornados.
Caso o bairro não exista, um erro é retornado.
*/

type FindNeighborhoodByIdUseCase struct {
	db *gorm.DB
}

func NewFindNeighborhoodByIdUseCase(db *gorm.DB) *FindNeighborhoodByIdUseCase {
	return &FindNeighborhoodByIdUseCase{db}
}

func (usecase *FindNeighborhoodByIdUseCase) Execute(id uint) (*domain.Neighborhood, error) {
	ctx := context.Background()

	neighborhood, err := gorm.G[domain.Neighborhood](usecase.db).Where("id = ?", id).First(ctx) 
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &neighborhood, nil
}
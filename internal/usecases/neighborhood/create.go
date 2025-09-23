package neighborhood

import (
	"codelab/backend/internal/domain"
	"context"

	"gorm.io/gorm"
)

/*
Caso de Uso:
Criar um novo bairro.

É passada o payload de dados para a criação do
novo bairro.
*/

type CreateNeighborhoodUseCase struct {
	db *gorm.DB
}

func NewCreateNeighborhoodUseCase(db *gorm.DB) *CreateNeighborhoodUseCase {
	return &CreateNeighborhoodUseCase{db}
}

func (usecase *CreateNeighborhoodUseCase) Execute(payload *domain.CreateNeighborhoodDTO) error {
	ctx := context.Background()

	neighborhood := domain.Neighborhood{
		Name: payload.Name,
		Latitude: payload.Latitude,
		Longitude: payload.Longitude,
		Habitants: payload.Habitants,
		Area: payload.Area,
		Active: payload.Active,
	}

	return gorm.G[domain.Neighborhood](usecase.db).Create(ctx, &neighborhood)
}
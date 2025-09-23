package neighborhood

import (
	"codelab/backend/internal/domain"
	"context"

	"gorm.io/gorm"
)

/*
Caso de Uso:
Buscar todos os bairros.
*/

type FindAllNeighborhoodsUseCase struct {
	db *gorm.DB
}

func NewFindAllNeighborhoodsUseCase(db *gorm.DB) *FindAllNeighborhoodsUseCase {
	return &FindAllNeighborhoodsUseCase{db}
}

func (usecase *FindAllNeighborhoodsUseCase) Execute() ([]domain.Neighborhood, error) {
	ctx := context.Background()

	return gorm.G[domain.Neighborhood](usecase.db).Find(ctx)
}
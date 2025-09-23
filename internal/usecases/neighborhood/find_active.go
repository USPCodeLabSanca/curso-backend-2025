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

type FindActiveNeighborhoodsUseCase struct {
	db *gorm.DB
}

func NewFindActiveNeighborhoodsUseCase(db *gorm.DB) *FindActiveNeighborhoodsUseCase {
	return &FindActiveNeighborhoodsUseCase{db}
}

func (usecase *FindActiveNeighborhoodsUseCase) Execute() ([]domain.Neighborhood, error) {
	ctx := context.Background()

	return gorm.G[domain.Neighborhood](usecase.db).Where("active = ?", true).Find(ctx)
}
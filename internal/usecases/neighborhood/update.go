package neighborhood

import (
	"codelab/backend/internal/domain"
	"context"

	"gorm.io/gorm"
)

/*
Caso de Uso:
Atualizar um bairro.

Os dados do bairro já atualizado
são informados e atualizados no banco.
*/

type UpdateNeighborhoodUseCase struct {
	db *gorm.DB
}

func NewUpdateNeighborhoodUseCase(db *gorm.DB) *UpdateNeighborhoodUseCase {
	return &UpdateNeighborhoodUseCase{db}
}

func (usecase *UpdateNeighborhoodUseCase) Execute(neighborhood *domain.Neighborhood) error {
	ctx := context.Background()

	_, err := gorm.G[domain.Neighborhood](usecase.db).Where("id = ?", neighborhood.ID).Updates(ctx, *neighborhood)

	return err
}
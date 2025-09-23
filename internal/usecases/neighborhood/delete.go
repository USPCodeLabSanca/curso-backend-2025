package neighborhood

import (
	"codelab/backend/internal/domain"
	"context"

	"gorm.io/gorm"
)

/*
Caos de Uso:
Remover bairro

O ID do bairro é especificado e o bairro
é removido.

Aqui, o que ocorre é um "soft-delete", ou seja,
os dados NÃO são deleteados de fato, apenas marcados como
deleteados.
Desse modo, dados marcados como deletados não podem
ser acessados via consultas (são ignorados).

Usando essa abordagem (padrão GORM), os dados de 
clima de um bairro permanecem lá, mesmo se o o bairro
estiver "deletado".
*/

type DeleteNeighborhoodUseCase struct {
	db *gorm.DB
}

func NewDeleteNeighborhoodUseCase(db *gorm.DB) *DeleteNeighborhoodUseCase {
	return &DeleteNeighborhoodUseCase{db}
}

func (usecase *DeleteNeighborhoodUseCase) Execute(id uint) error {
	ctx := context.Background()

	_, err := gorm.G[domain.Neighborhood](usecase.db).Where("id = ?", id).Delete(ctx)

	return err
}
package neighborhood

import (
	"codelab/backend/internal/domain"
	"codelab/backend/pkg/dates"
	"time"

	"gorm.io/gorm"
)

/*
Caso de Uso:
Buscar dados climáticos de um bairro
entre dois intervalos de tempo.
*/

type FindNeighborhoodWeatherUseCase struct {
	db *gorm.DB
}

func NewFindNeighborhoodWeatherUseCase(db *gorm.DB) *FindNeighborhoodWeatherUseCase {
	return &FindNeighborhoodWeatherUseCase{db}
}

/*
Aqui, a query do GORM é um pouco diferente
por ser da versão mais antiga.

A ideia é a mesma, apenas a interface de acesso mudou.

Nessa consulta, um JOIN é usado!
Isso significa:

Buscar na tabela de bairros o bairro com id = idBusca.
Com o bairro encontrado, juntar com a tabela de climas
e buscar todos os climas com idBairro = idBusca.
Além disso, ainda é necessário filtrar os climas por data,
então adiciona-se uma cláusula "WHERE" com as datas.

Em SQL seria algo como:
SELECT b.*, c.* from bairro b JOIN climas c ON b.id = c.idBairro WHERE b.id = ? AND (c.collected_at BETWEEN ? AND ?)

O GORM facilita consultas como essa, evitando escrever queries grandes
e com múltiplos joins.
*/
func (usecase *FindNeighborhoodWeatherUseCase) Execute(id uint, start, end string) (*domain.Neighborhood, error) {
	var neighborhood domain.Neighborhood

	// Parse do start
	startDate, err := dates.ParseDate(start)
	if err != nil {
		return nil, err
	}

	// Parse do end
	endDate, err := dates.ParseDate(end)
	if err != nil {
		return nil, err
	}

	endDate = endDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	// Montar a query
	err = usecase.db.Model(&domain.Neighborhood{}).
		Preload("Weather", "collected_at BETWEEN ? AND ?", startDate, endDate).
		First(&neighborhood, id).Error
		
	if err !=nil {
		return nil, err
	}

	return &neighborhood, nil
}
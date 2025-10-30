package domain

import (
	"time"
)

/*
Entidade "Bairro", definida usando o modelo
padrão do GORM e os atributos.

Aqui, a lista de "WeatherData" indica que um bairro
pode ter vários dados climáticos associados à ele.
Essa é uma relação "One to Many".

No banco de dados, a tabela "Bairro" NÃO
armazena essa lista!
É a tabela "Dados climáticos" que, em cada registro,
indica de qual bairro aquele dado está relacionado.

Veja "weather.go" para mais informações.
*/
type Neighborhood struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at,omitempty"`
	Name      string        `json:"name"`
	Latitude  float32       `json:"latitude"`
	Longitude float32       `json:"longitude"`
	Habitants uint32        `json:"habitants"`
	Area      float32       `json:"area"`
	Active    bool          `json:"active"`
	Weather   []WeatherData `json:"weather" gorm:"foreignKey:CollectedFrom;references:ID"`
}

/*
DTO para transferência de dados
usados na coleta da API.
*/
type NeighborhoodInfoDTO struct {
	ID  uint
	URL string
}

/*
DTO usado para receber a requisição
de cadastro de um novo bairro.
*/
type CreateNeighborhoodDTO struct {
	Name      string  `json:"name" binding:"required"`
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
	Habitants uint32  `json:"habitants" binding:"required"`
	Area      float32 `json:"area" binding:"required"`
	Active    bool    `json:"active"`
}

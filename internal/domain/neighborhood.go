package domain

import "gorm.io/gorm"

/*
Entidade "Bairro", definida usando o modelo
padrão do GORM e os atributos.
*/
type Neighborhood struct {
	gorm.Model
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Habitants uint32  `json:"habitants"`
	Area      float32 `json:"area"`
	Active    bool    `json:"active"`
}

/*
DTO para transferência de dados
usados na coleta da API.
*/
type NeighborhoodInfoDTO struct {
	ID int64
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
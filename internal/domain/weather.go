package domain

import (
	"time"

	"gorm.io/gorm"
)

/*
Definições dos códigos de clima e suas respectivas
mensagens em string.
A tabela está disponível no site da API, e foi
adaptada para código usando IA.
*/
type WeatherCode int

const (
	ClearSky WeatherCode = 0

	PartlyCloudy WeatherCode = 1
	MostlyCloudy WeatherCode = 2
	Overcast     WeatherCode = 3

	Mist         WeatherCode = 45
	FreezingMist WeatherCode = 48

	DrizzleLight    WeatherCode = 51
	DrizzleModerate WeatherCode = 53
	DrizzleHeavy    WeatherCode = 55

	FreezingDrizzleLight WeatherCode = 56
	FreezingDrizzleHeavy WeatherCode = 57

	RainLight    WeatherCode = 61
	RainModerate WeatherCode = 63
	RainHeavy    WeatherCode = 65

	FreezingRainLight WeatherCode = 66
	FreezingRainHeavy WeatherCode = 67

	SnowLight    WeatherCode = 71
	SnowModerate WeatherCode = 73
	SnowHeavy    WeatherCode = 75

	SnowGrains WeatherCode = 77

	ShowersLight    WeatherCode = 80
	ShowersModerate WeatherCode = 81
	ShowersViolent  WeatherCode = 82

	SnowShowersLight WeatherCode = 85
	SnowShowersHeavy WeatherCode = 86

	ThunderstormLight WeatherCode = 95

	ThunderstormHailLight WeatherCode = 96
	ThunderstormHailHeavy WeatherCode = 99
)

var WeatherDescriptions = map[WeatherCode]string{
	ClearSky: "Céu limpo",

	PartlyCloudy: "Parcialmente limpo",
	MostlyCloudy: "Parcialmente nublado",
	Overcast:     "Nublado",

	Mist:         "Névoa",
	FreezingMist: "Névoa com formação de gelo",

	DrizzleLight:    "Garoa leve",
	DrizzleModerate: "Garoa moderada",
	DrizzleHeavy:    "Garoa intensa",

	FreezingDrizzleLight: "Garoa congelante leve",
	FreezingDrizzleHeavy: "Garoa congelante intensa",

	RainLight:    "Chuva leve",
	RainModerate: "Chuva moderada",
	RainHeavy:    "Chuva forte",

	FreezingRainLight: "Chuva congelante leve",
	FreezingRainHeavy: "Chuva congelante forte",

	SnowLight:    "Neve leve",
	SnowModerate: "Neve moderada",
	SnowHeavy:    "Neve forte",

	SnowGrains: "Grãos de neve",

	ShowersLight:    "Pancadas de chuva leve",
	ShowersModerate: "Pancadas de chuva moderada",
	ShowersViolent:  "Pancadas de chuva violenta",

	SnowShowersLight: "Pancadas de neve leve",
	SnowShowersHeavy: "Pancadas de neve forte",

	ThunderstormLight: "Tempestade leve ou moderada",

	ThunderstormHailLight: "Tempestade com granizo leve",
	ThunderstormHailHeavy: "Tempestade com granizo forte",
}

/*
Entidade "Dados Climáticos", definida usando o modelo
padrão do GORM e os atributos.

Aqui, o campo "CollectedFrom" é uma chave estrangeria, ou seja,
é uma chave que vem de outra tabela do banco.

Nesse caso, "CollectedFrom" recebe o ID de algum bairro,
do qual este dado climático é referente.

Para buscar todos os dados climáticos de um bairro,
basta buscar todos os bairros com CollectedFrom = IdBairro.

Veja "neighborhood.go" para mais informações.
*/
type WeatherData struct {
	gorm.Model
	Temperature     float32     `json:"temperature"`
	Humidity        float32     `json:"humidity"`
	RainProbability float32     `json:"rainProbability"`
	RainVolume      float32     `json:"rainVolume"`
	Code            WeatherCode `json:"code"`
	CollectedAt     time.Time   `json:"collectedAt"`
	CollectedFrom   uint        `json:"collectedFrom"` //Neighborhood
}

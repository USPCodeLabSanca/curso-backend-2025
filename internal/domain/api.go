package domain

/*
Definição do modelo de dados retornado pela API.

Aqui, o modelo é simplificado, já que não precisamos
de todos os dados retornados.

Ao criar uma struct de resposta ('ApiResp') condizente
com o retorno da API, permitimos que a biblioteca de JSON
seja usada para converter a resposta diretamente para
nossa struct, o que evita um conversão manual, geralmente
campo a campo.
*/

type ApiCurrent struct {
	Temperature2m      float32 `json:"temperature_2m"`
	RelativeHumidity2m float32 `json:"relative_humidity_2m"`
	Rain               float32 `json:"rain"`
	WeatherCode        int     `json:"code"`
	Time               string  `json:"time"`
	Code               int     `json:"weather_code"`
}

type ApiResp struct {
	Current ApiCurrent `json:"current"`
}

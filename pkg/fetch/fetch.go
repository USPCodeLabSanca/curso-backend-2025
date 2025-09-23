package fetch

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
Função de fetch genérica.
Essa função só funciona utilizando GET.

Os dados são buscados na url e retornados
no formato de um map.
*/
func Fetch(url string) (map[string]any, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)

	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	
	return data, nil
}
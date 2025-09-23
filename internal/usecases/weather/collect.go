package weather

import (
	"codelab/backend/internal/domain"
	"codelab/backend/pkg/config"
	"codelab/backend/pkg/fetch"
	"context"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
)

/*
Caso de Uso:
Coletar as informações climáticas de todos os
bairros cadastrados que estão ativos.

Após criar a lista de bairros,
o gerenciador do processamento divide a carga de
trabalho criando uma "worker pool": conjunto de
goroutines que irão processar os dados.

Nesse modelo, definimos um número máximo de workers
(goroutines), e a lista de jobs é dividida entre eles.

Os resultados e erros são coletados via channels.
*/

type CollectWeatherDataUseCase struct {
	db *gorm.DB
	workers          int
	api              string
}

func NewCollectWeatherDataUseCase(
	db *gorm.DB,
	config *config.Config,
) *CollectWeatherDataUseCase {
	return &CollectWeatherDataUseCase{db, config.Workers, config.API}
}

func (usecase *CollectWeatherDataUseCase) Execute() error {
	ctx := context.Background()

	// Active neighborhoods
	neighborhoods, err :=gorm.G[domain.Neighborhood](usecase.db).Where("active = ?", "true").Find(ctx)
	if err != nil {
		return err
	}

	if len(neighborhoods) == 0 {
		return fmt.Errorf("no neighborhoods to fetch")
	}

	// Create target urls
	targets := make(chan domain.NeighborhoodInfoDTO)

	// Create channels
	errors := make(chan error)
	data := make(chan domain.WeatherData)

	for _, neighborhood := range neighborhoods {
		targets <- domain.NeighborhoodInfoDTO{
			ID: neighborhood.ID,
			URL: fmt.Sprintf("%v&latitude=%v&longitude=%v", usecase.api, neighborhood.Latitude, neighborhood.Longitude),
		}
		
	}

	var wg sync.WaitGroup
	wg.Add(int(usecase.workers))

	// Execute workers to fetch data
	for range usecase.workers {
		go worker(targets, errors, data, &wg)
	}

	wg.Wait()

	return nil
}

/*
Worker genérico para o processamento da API de clima.

O worker é responsável por pegar um job da lista e 
fazer a chamada na API para a URL especificada.
Cada worker cuida de algumas URLs e salva os
resultados no channel de dados.
*/
func worker(targets chan domain.NeighborhoodInfoDTO, errors chan error, data chan domain.WeatherData, wg *sync.WaitGroup) {
	defer wg.Done()

	for target := range targets {
		result, err := fetch.Fetch(target.URL)
		if err != nil {
			errors <- err
			continue
		}

		// Parse result and save
		data <- parse(result, target.ID)
	}
}

/*
Função responsável por converter os dados retornados
da API para o modelo definido em WeatherData.
*/
func parse(data map[string]any, neighborhood uint) domain.WeatherData {
	current, ok := data["current"].(map[string]any)
	if !ok {
		return domain.WeatherData{}
	}

	temp, ok := current["temperature_2m"].(float32)
	if !ok {
		temp = 0
	}

	humidity, ok := current["humidity"].(float32)
	if !ok {
		humidity = 0
	}

	rainProb, ok := current["rain_probability"].(float32)
	if !ok {
		rainProb = -1
	}

	return domain.WeatherData{
		Temperature: temp,
		Humidity: humidity,
		RainProbability: rainProb,
		CollectedFrom: neighborhood,
		CollectedAt: time.Now(),
	}
}

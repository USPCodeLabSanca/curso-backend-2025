package weather

import (
	"codelab/backend/internal/domain"
	"codelab/backend/pkg/config"
	"codelab/backend/pkg/fetch"
	"context"
	"fmt"
	"log"
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
	db      *gorm.DB
	workers int
	api     string
}

func NewCollectWeatherDataUseCase(
	db *gorm.DB,
	config *config.Config,
) *CollectWeatherDataUseCase {
	return &CollectWeatherDataUseCase{db, config.Workers, config.API}
}

func (usecase *CollectWeatherDataUseCase) Execute() error {
    ctx := context.Background()

    // Buscar bairros ativos
    neighborhoods, err := gorm.G[domain.Neighborhood](usecase.db).
        Where("active = ?", true).
        Find(ctx)

    if err != nil {
        return err
    }

    if len(neighborhoods) == 0 {
        return fmt.Errorf("no neighborhoods to fetch")
    }

    // Criar canais
    targets := make(chan domain.NeighborhoodInfoDTO, len(neighborhoods))
    errors := make(chan error, len(neighborhoods))
    data := make(chan domain.WeatherData, len(neighborhoods))

    // Enviar todos os bairros para o canal
    for _, neighborhood := range neighborhoods {
            targets <- domain.NeighborhoodInfoDTO{
                ID:  neighborhood.ID,
                URL: fmt.Sprintf("%v&latitude=%v&longitude=%v",
                    usecase.api, neighborhood.Latitude, neighborhood.Longitude),
            }
        }
        close(targets)

    // Workers
    var wg sync.WaitGroup

    for i := 0; i < usecase.workers; i++ {
		wg.Add(1)
        go worker(targets, errors, data, &wg)
    }

    wg.Wait()
    close(data)
    close(errors)

    // Consumir e salvar resultados
    for weather := range data {
        if err := usecase.db.WithContext(ctx).Create(&weather).Error; err != nil {
            log.Printf("error saving weather data: %v", err)
        }
    }

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
        log.Printf("Fetching weather data for neighborhood %v", target.ID)

        result, err := fetch.Fetch(target.URL)
        if err != nil {
            errors <- fmt.Errorf("failed to fetch from %s: %w", target.URL, err)
            continue
        }

       	weather, err := parse(result, target.ID)
        if err != nil {
            errors <- fmt.Errorf("failed to parse weather for neighborhood %d: %w", target.ID, err)
            continue
        }

        data <- weather
    }
}

/*
Função responsável por converter os dados retornados
da API para o modelo definido em WeatherData.
*/
func parse(data *domain.ApiResp, neighborhood uint) (domain.WeatherData, error) {
    return domain.WeatherData{
        Temperature:     data.Current.Temperature2m,
        Humidity:        data.Current.RelativeHumidity2m,
        RainProbability: data.Current.Rain,
        Code:            domain.WeatherCode(data.Current.Code),
        CollectedFrom:   neighborhood,
        CollectedAt:     time.Now(),
        CreatedAt:       time.Now(),
        UpdatedAt:       time.Now(),
    }, nil
}

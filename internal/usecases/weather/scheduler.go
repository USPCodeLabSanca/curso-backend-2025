package weather

import (
	"log"
	"time"
)

/*
Sistema que agenda a coleta de dados em cada bairro.
Ao inicarmos o servidor, o scheduler também é iniciado.

O valor definido em 'interval' define o tempo entre cada coleta.
O scheduler chama o caso de uso "coletar dados climáticos" sempre que
necessário.

Ao encerrar o servidor, o scheduler também é parado
*/

type WeatherScheduler struct {
	usecase  *CollectWeatherDataUseCase
	interval time.Duration
}

// Inicialização com o caos de uso CollectWeatherData
func NewWeatherScheduler(usecase *CollectWeatherDataUseCase) *WeatherScheduler {
	return &WeatherScheduler{
		usecase:  usecase,
		interval: time.Minute,
	}
}

// Inicia a goroutine do ticker, que conta o intervalo
func (s *WeatherScheduler) Start(stopChan <-chan struct{}) {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Collecting weather data")
			if err := s.usecase.Execute(); err != nil {
				log.Println("Failed to collect data:", err)
			}
		case <-stopChan:
			log.Println("Weather scheduler stopped.")
			return
		}
	}
}

// Função para iniciar um scheduler e colocá-lo para executar.
func InitWeatherScheduler(scheduler *WeatherScheduler) chan struct{} {
	stopChan := make(chan struct{})

	go func() {
		log.Println("Scheduler started")
		scheduler.Start(stopChan)
		log.Println("Scheduler stopped")
	}()

	return stopChan
}

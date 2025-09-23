package weather

import "go.uber.org/fx"

var Module = fx.Module("weather-usecases",
	fx.Provide(
		NewCollectWeatherDataUseCase,
	),
)
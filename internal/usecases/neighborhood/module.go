package neighborhood

import "go.uber.org/fx"

var Module = fx.Module("neighborhood-usecases",
	fx.Provide(
		NewCreateNeighborhoodUseCase,
		NewDeleteNeighborhoodUseCase,
		NewFindActiveNeighborhoodsUseCase,
		NewFindAllNeighborhoodsUseCase,
		NewFindNeighborhoodByIdUseCase,
		NewUpdateNeighborhoodUseCase,
	),
)
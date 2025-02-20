package author

import (
	"go.uber.org/fx"
)

var Module = fx.Module("author",
	fx.Options(
		fx.Provide(
			NewService,
			NewController,
			NewRepository,
			NewRoute,
		),
		fx.Invoke(RegisterRoute),
	),
)

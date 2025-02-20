package seeds

import (
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"

	"go.uber.org/fx"
)

func AsSeeder(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(framework.Seed)),
		fx.ResultTags(`group:"seeds"`),
	)
}

var (
	Module = fx.Module("seeds",
		fx.Provide(
			AsSeeder(NewHelloSeed),
			fx.Annotate(
				NewSeeder,
				fx.ParamTags(`group:"seeds"`),
			),
		),
	)
)

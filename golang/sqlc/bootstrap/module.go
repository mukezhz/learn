package bootstrap

import (
	"github.com/mukezhz/learn/tree/main/golang/sqlc/internal/domain"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/migrations"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/seeds"

	"go.uber.org/fx"
)

var CommonModules = fx.Module("common",
	fx.Options(
		pkg.Module,
		seeds.Module,
		migrations.Module,
		domain.Module,
	),
)

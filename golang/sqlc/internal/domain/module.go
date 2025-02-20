package domain

import (
	"github.com/mukezhz/learn/tree/main/golang/sqlc/internal/domain/author"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/internal/domain/middlewares"
	"go.uber.org/fx"
)

var Module = fx.Options(
	middlewares.Module,
	author.Module,
)

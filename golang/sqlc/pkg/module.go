package pkg

import (
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/infrastructure"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/middlewares"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/services"

	"go.uber.org/fx"
)

var Module = fx.Module("pkg",
	framework.Module,
	services.Module,
	middlewares.Module,
	infrastructure.Module,
)

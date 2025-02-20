package seeds

import (
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
)


type HelloSeed struct {
	logger         framework.Logger
	env            *framework.Env
}


func NewHelloSeed(
	logger framework.Logger,
	env *framework.Env,
) *HelloSeed {
	return &HelloSeed{
		logger:         logger,
		env:            env,
	}
}

func (as *HelloSeed) Name() string {
	return "HelloSeed"
}


func (as *HelloSeed) Seed() error {
	// inject the required service or repository and seed the data
	as.logger.Infoln("[Seeding...] seeding data for hello")
	return nil
}

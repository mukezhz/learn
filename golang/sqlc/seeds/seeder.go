package seeds

import (
	"errors"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
)

type Seeder struct {
	seeds  []framework.Seed
	logger framework.Logger
}

func NewSeeder(
	seeds []framework.Seed,
	logger framework.Logger,
) *Seeder {
	return &Seeder{
		seeds:  seeds,
		logger: logger,
	}
}

func (s *Seeder) Exec(names []string, runAll bool) error {
	s.logger.Infoln("🌱 seeding data...")
	if runAll {
		for _, seed := range s.seeds {
			err := seed.Seed()
			if err != nil {
				return err
			}
		}
		return nil
	}

	if len(names) == 0 {
		s.logger.Info("no seed name provided")
		return nil
	}

	for _, name := range names {
		if err := s.runSeed(name); err != nil {
			s.logger.Infof("Error running %s: %s", name, err)
		}
	}
	return nil
}

func (s *Seeder) runSeed(name string) error {
	isValid := false
	for _, seed := range s.seeds {
		if name == seed.Name() {
			isValid = true
			if err := seed.Seed(); err != nil {
				return err
			}
		}
	}

	if !isValid {
		return errors.New("invalid seed name")
	}
	return nil
}

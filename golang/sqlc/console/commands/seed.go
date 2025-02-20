package commands

import (
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/seeds"

	"github.com/spf13/cobra"
)

type SeedCommand struct {
	names  []string
	runAll bool
}

func (s *SeedCommand) Short() string {
	return "run seed command"
}

func NewSeedCommand() *SeedCommand {
	return &SeedCommand{}
}

func (s *SeedCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().StringArrayVarP(
		&s.names,
		"name",
		"n",
		[]string{},
		"name of the seed to run (can be used multiple times)",
	)
	cmd.Flags().BoolVar(&s.runAll, "all", false, "run all seeds")
}

func (s *SeedCommand) Run() framework.CommandRunner {
	return func(
		l framework.Logger,
		seeder *seeds.Seeder,
	) {
		l.Info("Running seed command", s.names, s.runAll)
		if s.runAll {
			if err := seeder.Exec(nil, true); err != nil {
				l.Fatal(err)
			}
			return
		}
		if err := seeder.Exec(s.names, false); err != nil {
			l.Fatal(err)
		}

	}
}

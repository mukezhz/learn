package console

import (
	"context"

	"github.com/mukezhz/learn/tree/main/golang/sqlc/console/commands"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var cmds = map[string]framework.Command{
	"cmd:random": commands.NewRandomCommand(),
	"app:serve":  commands.NewServeCommand(),
	"seed:run":   commands.NewSeedCommand(),
	"migrate":    commands.NewMigrateCommand(),
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(opt fx.Option) []*cobra.Command {
	subCommands := make([]*cobra.Command, 0)
	for name, cmd := range cmds {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}

func WrapSubCommand(name string, cmd framework.Command, opt fx.Option) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
			logger := framework.GetLogger()

			opts := fx.Options(
				fx.WithLogger(logger.GetFxLogger),
				fx.Invoke(cmd.Run()),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer func() {
				err = app.Stop(ctx)
				if err != nil {
					logger.Fatal(err)
				}
			}()
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	cmd.Setup(wrappedCmd)
	return wrappedCmd
}

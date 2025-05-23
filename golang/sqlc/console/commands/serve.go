package commands

import (
    "time"
    "github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
    "github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/infrastructure"
    "github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/middlewares"

    "github.com/getsentry/sentry-go"
    "github.com/spf13/cobra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
    return "serve application"
}

func (s *ServeCommand) Setup(_ *cobra.Command) {}

func (s *ServeCommand) Run() framework.CommandRunner {
    return func(
            middleware middlewares.Middlewares,
            env *framework.Env,
            router *infrastructure.Router,
            logger framework.Logger,
    ) {
        logger.Info(`+-----------------------+`)
        logger.Info(`| GO CLEAN ARCHITECTURE |`)
        logger.Info(`+-----------------------+`)

        // Using time zone as specified in env file
        loc, _ := time.LoadLocation(env.TimeZone)
        time.Local = loc

        middleware.Setup()

        if env.Environment != "local" && env.SentryDSN != "" {
            err := sentry.Init(sentry.ClientOptions{
                Dsn:              env.SentryDSN,
                AttachStacktrace: true,
            })
            if err != nil {
                logger.Error("sentry initialization failed")
                logger.Error(err.Error())
            }
        }
        logger.Info("Running server")
        if env.ServerPort == "" {
            if err := router.Run(); err != nil {
                logger.Fatal(err)
                return
            }
        } else {
            if err := router.Run(":" + env.ServerPort); err != nil {
                logger.Fatal(err)
                return
            }
        }
    }
}

func NewServeCommand() *ServeCommand {
    return &ServeCommand{}
}

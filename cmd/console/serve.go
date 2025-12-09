package console

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/newbie007fx/cinemas/internal/dependencies"
	"github.com/newbie007fx/cinemas/internal/transport/http/routes"
	"github.com/newbie007fx/cinemas/platform/configuration"
	"github.com/newbie007fx/cinemas/platform/database"
	"github.com/newbie007fx/cinemas/platform/httpserver"
	"github.com/newbie007fx/cinemas/platform/validation"

	"github.com/spf13/cobra"
)

func (cs Console) initServeCommand() {
	cmd := cs.ConsoleService.GetCommandInstance()

	cmd.Use = "serve"

	cmd.Short = "Run service"

	cmd.Run = func(_ *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		config := configuration.New("./config", "config", "yaml")
		config.Setup()

		logLevel := &slog.LevelVar{}
		logLevel.UnmarshalText([]byte(config.GetConfig().App.LogLevel))
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		}))
		slog.SetDefault(logger)

		db := database.New(config)
		db.Setup()

		validator := validation.New()
		validator.Setup()

		server := httpserver.New(config)
		server.Setup()

		dep := dependencies.New(db, config)
		dep.Init()

		routes.Init(server, dep)

		server.Start()

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-done
		log.Println("shutting down")

		server.Shutdown(ctx)
		db.Shutdown()

		log.Println("all server stopped!")
	}

	cs.ConsoleService.RegisterCommand(cmd)
}

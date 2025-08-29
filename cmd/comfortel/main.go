package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	application "github.com/srgklmv/comfortel/internal/app"
	"github.com/srgklmv/comfortel/pkg/logger"
)

func main() {
	logger.Init()
	logger.Info("Starting up Comfortel...")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	app := application.New()

	go func() {
		if err := app.Run(); err != nil {
			logger.Error("Startup error. Exiting Comfortel...", slog.String("error", err.Error()))
			shutdown <- syscall.SIGTERM
		}
	}()
	logger.Info("Comfortel is running.")

	<-shutdown
	logger.Info("Shutting down Comfortel...")

	app.Shutdown()
	logger.Info("Comfortel shut down.")
}

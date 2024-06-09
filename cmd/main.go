package main

import (
	"os"
	"os/signal"
	"syscall"

	"dealls-dating-app/internal/provider"

	"go.uber.org/zap"
)

func main() {
	logger := provider.LoggerProvider()
	defer logger.Sync()

	app := provider.ProvideApp()

	go func() {
		if err := app.Run(); err != nil {
			logger.Fatal("app failed to run", zap.Error(err))
		}
	}()

	// listen to interrupt/terminate signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("terminate called")

	// call gracefully shutdown
	if err := app.Stop(); err != nil {
		logger.Fatal("error when shutting down app", zap.Error(err))
	}

	logger.Info("good bye.")
}

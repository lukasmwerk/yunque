package main

import (
	"os"

	"github.com/lukasmwerk/yunque/libs/logger"
	"github.com/lukasmwerk/yunque/services/core/config"
	"github.com/lukasmwerk/yunque/services/core/server"
)

func main() {
	ff := config.NewFeatureFlags()
	if !ff.IsValid() {
		panic("Feature flags invalid, cannot start service")
	}

	logger := logger.NewLogger(ff.LogMode)
	logger.UpdateStatus("Starting Server...")
	session := server.NewSession(ff.RunMode, logger)
	logger.UpdateStatus("Running Server...")
	session.Run() // enter run loop
	logger.UpdateStatus("Shutting down gracefully...")

	os.Exit(session.ExitCode)
}

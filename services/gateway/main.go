package main

import (
	"os"

	"github.com/lukasmwerk/yunque/libraries/logger"
	"github.com/lukasmwerk/yunque/services/gateway/config"
	"github.com/lukasmwerk/yunque/services/gateway/server"
)

func main() {
	ff := config.NewFeatureFlags()
	if !ff.IsValid() {
		panic("Feature flags invalid, cannot start service")
	}

	logger := logger.NewLogger(ff.LogMode)
	logger.UpdateStatus("Starting gateway service server...")
	session := server.NewSession(logger)
	logger.UpdateStatus("Running gateway service server...")
	session.Run() // enter run loop
	logger.UpdateStatus("Shutting down gracefully...")

	os.Exit(session.ExitCode)
}

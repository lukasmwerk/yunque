package main

import (
	"os"

	"github.com/lukasmwerk/yunque/libraries/logger"
	"github.com/lukasmwerk/yunque/services/inventory/server"
)

func main() {
	// ff := config.NewFeatureFlags()
	// if !ff.IsValid() {
	// 	panic("Feature flags invalid, cannot start service")
	// }

	logger := logger.NewLogger(0)
	logger.UpdateStatus("Starting inventory service server...")
	session := server.NewSession(logger)
	logger.UpdateStatus("Running inventory service server...")
	session.Run() // enter run loop
	logger.UpdateStatus("Shutting down gracefully...")

	os.Exit(session.ExitCode)
}

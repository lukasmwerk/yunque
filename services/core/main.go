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
		return
	}

	logger := logger.NewLogger(ff.LogMode)
	logger.UpdateStatus("Starting Server...")
	session := server.NewSession(ff.RunMode, logger)
	session.Run() // enter run loop
	logger.UpdateStatus("Shutting down gracefully...")

	os.Exit(session.ExitCode)
}

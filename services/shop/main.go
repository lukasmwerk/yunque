package main

import (
	"github.com/lukasmwerk/yunque/libraries/logger"
	"github.com/lukasmwerk/yunque/services/shop/server"
	"os"
)

func main() {
	ff := server.NewFeatureFlags()
	if !ff.IsValid() {
		panic("Feature flags invalid, cannot start service")
	}

	logger := logger.NewLogger(ff.LogMode)
	logger.UpdateStatus("Starting shop service server...")

	session, err := server.NewSession(logger)
	if err != nil {
		panic(err)
	}

	logger.UpdateStatus("Running shop service server...")
	err = session.Run() // enter run loop
	if err != nil {
		panic(err)
	}

	logger.UpdateStatus("Shutting down gracefully...")
	os.Exit(session.ExitCode)
}

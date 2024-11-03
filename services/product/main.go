package main

import (
	"os"

	"github.com/lukasmwerk/yunque/libraries/logger"
	"github.com/lukasmwerk/yunque/services/product/server"
)

func main() {
	logger := logger.NewLogger(0)
	logger.UpdateStatus("Starting product service server...")
	session := server.NewSession(logger)
	logger.UpdateStatus("Running product service server...")
	session.Run() // enter run loop
	logger.UpdateStatus("Shutting down gracefully...")

	os.Exit(session.ExitCode)
}

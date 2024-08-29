package main

import (
	"os"

	"github.com/lukasmwerk/yunque/libs/logger"
)

func main() {
	ff := NewFeatureFlags()
	if !ff.IsValid() {
		return
	}

	logger := logger.NewLogger(ff.LogMode)
	logger.UpdateStatus("Starting Server...")
	session := NewSession(ff.RunMode, logger)
	session.Run() // enter run loop
	logger.UpdateStatus("Shutting down gracefully")

	os.Exit(session.ExitCode)
}

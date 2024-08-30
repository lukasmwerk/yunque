package api

import (
	"net/http"

	"github.com/lukasmwerk/yunque/libs/logger"
)

// ConfigureServer sets up the http server that will handle the api
func ConfigureServer(logger logger.Logger) *http.Server {
	multiplexer := http.NewServeMux()

	// multiplexer.HandleFunc("/api/task", HandleTaskRequest)
	// multiplexer.HandleFunc("/api/status", HandleStatusRequest)

	server := &http.Server{
		Addr:    ":8080", // Port to listen on
		Handler: multiplexer,
	}

	logger.UpdateConfig("core rest api http server using port 8080...")
	logger.UpdateStatus("Starting server on port 8080...")

	return server
}

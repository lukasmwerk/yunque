package api

import (
	"fmt"
	"log"
	"net"

	"github.com/lukasmwerk/yunque/libs/logger"
	"google.golang.org/grpc"
)

// ConfigureServer sets up the grpc server that will handle communication
func ConfigureServer(logger logger.Logger) (*grpc.Server, *net.Listener) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8070))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	server := grpc.NewServer()

	logger.UpdateConfig("inventory grpc server listening on port 8070...")
	logger.UpdateStatus("Starting server on port 8070...")

	return server, &listener
}

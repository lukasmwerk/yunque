package server

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/lukasmwerk/yunque/libraries/logger"
	"github.com/lukasmwerk/yunque/libraries/types"
	"github.com/lukasmwerk/yunque/services/inventory/api"
	"google.golang.org/grpc"
)

type Session struct {
	State    types.RunState
	Context  context.Context
	ExitCode int
	Logger   logger.Logger
	Server   *grpc.Server
	Listener *net.Listener
	Services types.ServiceList
	Mutex    types.RWMutex
}

func NewSession(logger logger.Logger) *Session {
	server, listener := api.ConfigureServer(logger)
	session := &Session{
		Logger:   logger,
		Server:   server,
		Listener: listener,
	}

	return session
}

// Configure abstracts configurations to enable changing config without restarting the server
func (s *Session) Configure() {}

func (s *Session) Run() {
	go func() {
		if err := s.Server.Serve(*s.Listener); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()
	for {
		// s.checkServiceHealth() different implementation needed
		time.Sleep(time.Second)
	}
}

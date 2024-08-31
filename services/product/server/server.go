package server

import (
	"context"
	"net/http"
	"time"

	"github.com/lukasmwerk/yunque/libs/logger"
	"github.com/lukasmwerk/yunque/libs/types"
	"github.com/lukasmwerk/yunque/services/core/config"
)

type Session struct {
	State    types.RunState
	Context  context.Context
	Config   config.SessionConfig
	ExitCode int
	Logger   logger.Logger
	Server   *http.Server
	Services types.ServiceList
	Mutex    types.RWMutex
}

func NewSession(logger logger.Logger) *Session {
	session := &Session{
		Logger: logger,
		// Server: api.ConfigureServer(logger),
	}

	return session
}

// Configure abstracts configurations to enable changing config without restarting the server
func (s *Session) Configure() {}

func (s *Session) Run() {
	for range 5 {
		// s.checkServiceHealth()
		time.Sleep(time.Second)
	}
}

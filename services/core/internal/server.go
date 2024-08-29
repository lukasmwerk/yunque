package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/lukasmwerk/yunque/libs/logger"
	"github.com/lukasmwerk/yunque/libs/types"
	"github.com/lukasmwerk/yunque/services/core/api"
)

type Session struct {
	CLI      bool
	State    types.RunState
	Context  context.Context
	ExitCode int
	Logger   logger.Logger
	Server   *http.Server
}

func NewSession(rs types.RunState, logger logger.Logger) *Session {
	session := &Session{
		CLI:    true,
		State:  rs,
		Logger: logger,
		Server: api.ConfigureServer(logger),
	}

	return session
}

// Configure abstracts configurations to enable changing config without restarting the server
func (s *Session) Configure() {}

func (s *Session) Run() {
	for {
		checkServiceHealth()
		time.Sleep(time.Second)
	}
}

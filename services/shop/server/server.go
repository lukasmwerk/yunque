package server

import (
	"github.com/lukasmwerk/yunque/libraries/logger"
)

// Session is a shop server session. In the short term this pattern allows
// spawning in many shop sessions for testing.
type Session struct {
	Logger   logger.Logger
	ExitCode int
}

// NewSession returns a shop server session that can be run.
func NewSession(logger logger.Logger) (*Session, error) {
	return &Session{
		Logger: logger,
	}, nil
}

// Run runs the created shop server session.
func (s *Session) Run() error {
	return nil
}

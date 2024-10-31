package config

// SessionConfig contains all relevant configurations for a gateway server session
type SessionConfig struct {
	ServiceNames []string // expected services to be connected
}

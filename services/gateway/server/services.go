package server

import "github.com/lukasmwerk/yunque/libraries/common"

// Checks status of registered services that are expected by the server configuration
func (s *Session) checkServiceHealth() error {
	if len(s.Config.ServiceNames) > len(s.Services) {
		defer s.resolveMissingServices()
	}
	return nil
}

// resolveMissingServices makes sure unregistered services are resolved
func (s *Session) resolveMissingServices() {
	for _, es := range s.Config.ServiceNames {
		var found bool
		for _, rs := range s.Services {
			if es == rs.Name {
				found = true
			}
		}
		if !found {
			_, err := common.Retry(s.RegisterService(es))
			if err != nil {
				// TODO: handle error
			}
		}
	}
}

func (s *Session) RegisterService(name string) error {
	return nil
}

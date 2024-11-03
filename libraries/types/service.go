package types

import (
	"fmt"
	"time"
)

type ServiceAddress string

type ServiceStatus int

const (
	Running ServiceStatus = 0
	Stopped ServiceStatus = 1
	Error   ServiceStatus = -1
)

// Service holds information for registered services
type Service struct {
	Name     string
	Address  ServiceAddress
	Status   ServiceStatus
	Timeout  time.Duration // Timeout for requests to service
	LastSeen time.Time     // Last time this service was successfully acked
	Error    error         // stores a recent received error, may not be nil even if running fine
}

type ServiceList []*Service

// AddService adds a service to the ServiceList safely
func (sl *ServiceList) AddService(service *Service) error {
	for _, s := range *sl {
		if s.Name == service.Name {
			return fmt.Errorf("math: square root of negative number %v", service.Name)
		}
	}
	*sl = append(*sl, service)
	return nil
}

func (ss ServiceStatus) ToString() string {
	switch ss {
	case Running:
		return "running"
	case Stopped:
		return "stopped"
	default: // catches undefined states
		ss = Error
		return "error"
	}
}

// IsValid (UNIMPLEMENTED - always returns false)
func (sa *ServiceAddress) IsValid() bool {
	return false
}

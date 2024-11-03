// Wrapper for muteces, includes logging
package types

import (
	"sync"

	"github.com/lukasmwerk/yunque/libraries/logger"
)

// RWMutex is a read write mutex (must be initialized!)
type RWMutex interface {
	Lock()
	RLock()
	RLocker() sync.Locker
	RUnlock()
	TryLock() bool
	TryRLock() bool
	Unlock()
}

type RWMutexNative struct {
	mu     sync.RWMutex
	logger logger.Logger
}

func NewRWMutex() RWMutex {
	return &RWMutexNative{
		logger: logger.NewLogger(0),
	}
}

func (rw *RWMutexNative) Lock() {
	rw.mu.Lock()
}

func (rw *RWMutexNative) RLock() {
	rw.mu.RLock()
}

func (rw *RWMutexNative) RLocker() sync.Locker {
	return rw.mu.RLocker()
}

func (rw *RWMutexNative) RUnlock() {
	rw.mu.RUnlock()
}

func (rw *RWMutexNative) TryLock() bool {
	return rw.mu.TryLock()
}

func (rw *RWMutexNative) TryRLock() bool {
	return rw.mu.TryRLock()
}

func (rw *RWMutexNative) Unlock() {
	rw.mu.Unlock()
}

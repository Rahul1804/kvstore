package store

import (
	"sync"

	"github.com/rahul1804/kvstore/pkg/logger"
)

// Store represents a thread-safe key-value store.
type Store struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewStore creates a new Store.
func NewStore() *Store {
	return &Store{
		store: make(map[string]string),
	}
}

// Get retrieves the value for a given key.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.store[key]
	return value, ok
}

// Set sets the value for a given key.
func (s *Store) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
	logger.Logger.Infof("Key set: %s = %s", key, value)
}

// Delete removes a key-value pair.
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
	logger.Logger.Infof("Key deleted: %s", key)
}

// GetAll returns all key-value pairs.
func (s *Store) GetAll() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// Create a copy to avoid race conditions
	copy := make(map[string]string)
	for key, value := range s.store {
		copy[key] = value
	}
	return copy
}

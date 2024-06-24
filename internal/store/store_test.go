package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	s := NewStore()

	s.Set("foo", "bar")
	value, ok := s.Get("foo")
	if !ok || value != "bar" {
		t.Errorf("Expected to get 'bar' for key 'foo', got '%s'", value)
	}

	s.Delete("foo")
	_, ok = s.Get("foo")
	if ok {
		t.Errorf("Expected key 'foo' to be deleted")
	}
}

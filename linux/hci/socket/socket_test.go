package socket

import (
	"testing"
)

// TestSocketIsClosed verifies the behavior of the IsClosed() method.
// It checks that IsClosed() returns false when the socket is open and
// true after the channel is closed.
func TestSocketIsClosed(t *testing.T) {
	// Create a new Socket which is initially not closed.
	s := &Socket{
		closed: make(chan struct{}),
	}

	// Verify that IsClosed() returns false for an open socket.
	if s.IsClosed() {
		t.Errorf("Expected socket to be open, but IsClosed() returned true")
	}

	// Close the channel to simulate closing the socket.
	close(s.closed)

	// After closing the channel, IsClosed() should return true.
	if !s.IsClosed() {
		t.Errorf("Expected socket to be closed, but IsClosed() returned false")
	}

	// Calling Close() again should have no effect and should not produce an error.
	if err := s.Close(); err != nil {
		t.Fatalf("Close called again returned error: %v", err)
	}

	// IsClosed() should still return true after attempting to close again.
	if !s.IsClosed() {
		t.Errorf("Expected socket to still be closed, but IsClosed() returned false")
	}

	// Reassign a new channel to simulate opening the socket again.
	s.closed = make(chan struct{})

	// Now IsClosed() should return false again for the "reopened" socket.
	if s.IsClosed() {
		t.Errorf("Expected socket to be open, but IsClosed() returned true")
	}
}

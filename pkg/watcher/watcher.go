package watcher

import (
	"errors"
	"io/fs"
)

// Op describes the type of file system operation.
type Op uint32

const (
	// Create indicates a new file or directory was created.
	Create Op = 1 << iota
	// Write indicates a file or directory was written to.
	Write
	// Rename indicates a file or directory was renamed.
	Rename
	// Remove indicates a file or directory was removed.
	Remove
)

// Event represents a file system event.
type Event struct {
	Path string      // Path to the file or directory.
	Op   Op          // Operation that triggered the event.
	Info fs.FileInfo // File metadata at the time of the event.
}

// Watcher defines the behavior of a file system watcher.
type Watcher interface {
	// Start begins watching for file system changes.
	Start() error
	// Close stops watching and releases any associated resources.
	Close() error
	// Events returns a channel that receives file system events.
	Events() <-chan Event
	// Errors returns a channel that receives watcher errors.
	Errors() <-chan error
}

// ErrUnsupported is returned when a watcher operation is not supported on the
// current platform.
var ErrUnsupported = errors.New("watcher: unsupported platform or operation")

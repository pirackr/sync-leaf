package watcher

import "testing"

func TestErrUnsupported(t *testing.T) {
	if ErrUnsupported == nil {
		t.Fatalf("ErrUnsupported is nil")
	}
}

func TestOpValues(t *testing.T) {
	ops := []Op{Create, Write, Rename, Remove}
	for _, op := range ops {
		if op == 0 {
			t.Fatalf("op %v is zero", op)
		}
	}
}

func TestEvent(t *testing.T) {
	e := Event{Path: "file.txt", Op: Create}
	if e.Path != "file.txt" {
		t.Fatalf("Event.Path = %q, want %q", e.Path, "file.txt")
	}
}

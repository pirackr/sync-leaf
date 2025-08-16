package main

import "testing"

func TestRun(t *testing.T) {
	want := "Hello, world!"
	if got := run(); got != want {
		t.Fatalf("run() = %q, want %q", got, want)
	}
}

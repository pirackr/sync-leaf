package greetings

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world!"
	if got := Hello(); got != want {
		t.Fatalf("Hello() = %q, want %q", got, want)
	}
}

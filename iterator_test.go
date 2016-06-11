package iterator

import (
	"testing"
)

func TestPing(t *testing.T) {
	r := Ping()
	if expected := "pong"; r != expected {
		t.Fatalf("expected %s, actual %s", expected, r)
	}
}

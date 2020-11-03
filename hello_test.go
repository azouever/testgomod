package testgomod

import "testing"

func TestHello(t *testing.T) {
	want := "hello, world."
	if got := Hello("world."); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nick")

	got := buffer.String()
	want := "Hello, Nick"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

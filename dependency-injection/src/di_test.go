package dependencyInjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Enrique")

	got := buffer.String()
	want := "Hello, Enrique"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

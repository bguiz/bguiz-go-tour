package main

import (
	"bytes"
	"testing"
)

func TestGreetNoDi(t *testing.T) {
	GreetNoDi("JJ")
}

// NOTE
// - but we can't actually test the output, because
//   we cannot intercept calls to fmt.Printf()
// - so drill down:
//   - `fmt.Printf`
//   - `fmt.Fprintf`
//   - `io.Writer` -> interface
// - thus we can test this by refactoring our code to
//   pass in an object that implements `io.Writer`

func TestGeetWithDi(t *testing.T) {
	buffer := bytes.Buffer{}
	GreetWithDi(&buffer, "JJ")

	got := buffer.String()
	want := "Hello JJ"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

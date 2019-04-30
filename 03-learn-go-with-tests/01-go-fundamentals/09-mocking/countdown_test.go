package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}

	Countdown(&buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestCountdownWithTime(t *testing.T) {
	buffer := bytes.Buffer{}

	CountdownWithTime(&buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(seconds int) {
	s.Calls++
}

// NOTE
// - we create an abstraction in the implementation, which specifies the
//   interface called Sleeper, and provide two implementations for it
//   one for use in the specifications (tests)
// - spies are a type of mock, which also keeps records of how it is used

func TestCountdownWithSleeper(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeper := SpySleeper{}

	CountdownWithSleeper(&buffer, &sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if sleeper.Calls != 3 {
		t.Errorf("expected sleeper to be called 3 times, was called %d times", sleeper.Calls)
	}
}

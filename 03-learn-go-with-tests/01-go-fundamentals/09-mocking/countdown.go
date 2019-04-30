package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

const FINAL_COUNTDOWN = "Go!"

func CountdownWithTime(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, FINAL_COUNTDOWN)
}

// NOTE
// - when we test this, we need to actually wait 3s
//   for the tests to execute, so let's mock it instead

type Sleeper interface {
	Sleep(seconds int)
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func CountdownWithSleeper(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		sleeper.Sleep(1)
	}
	fmt.Fprint(out, FINAL_COUNTDOWN)
}

func main() {
	// CountdownWithTime(os.Stdout)
	// CountdownWithTime(os.Stdout, time) // ERROR use of package time without selector
	sleeper := DefaultSleeper{}
	CountdownWithSleeper(os.Stdout, &sleeper)
}

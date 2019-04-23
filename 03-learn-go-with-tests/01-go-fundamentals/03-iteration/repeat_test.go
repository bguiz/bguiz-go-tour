package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func BenchmarkRepeat1Char5Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeat1Char20Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 20)
	}
}

func BenchmarkRepeat26Char5Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("abcdefghijklmnopqrstuvwxyz", 5)
	}
}

func BenchmarkRepeat26Char20Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("abcdefghijklmnopqrstuvwxyz", 20)
	}
}

func BenchmarkStringsRepeat26Char20Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// https://golang.org/pkg/strings/#Repeat
		strings.Repeat("abcdefghijklmnopqrstuvwxyz", 20)
	}
}

// BenchmarkRepeat26Char20Times-4          	 1000000	      1527 ns/op
// BenchmarkStringsRepeat26Char20Times-4   	10000000	       139 ns/op
// -> 1527 vs 139 ns/op - clearly using the core go implementation yields superior performance!

// NOTE
// - you can run benchmarks like testing, using `go test -bench=.`
// - the function name must start with `Benchmark`, and it takes in a `*testing.B`
// - use a for loop to run the function that you wish to run a specified number of times

func ExampleRepeat() {
	result := Repeat("Abc", 3)
	fmt.Println(result)
	// Output: AbcAbcAbc
}

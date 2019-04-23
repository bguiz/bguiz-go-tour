package main

import "testing"

/*
func DeprecatedTestHello(t *testing.T) {
	got := Hello("bguiz")
	want := "Hello, bguiz"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
*/

// NOTE
// - run tests using `go test`
// - if you have installed `github.com/mdempsky/gocode` and `github.com/go-delve/delve/cmd/dlv`,
//   and you are using an IDE that integrates with them, e.g. VS Code,
//   you should be able to run/ debug the code within the IDE itself
// - test file must be named `${something}_test.go`, test function name must start with `Test`
// - `*testing.T` is the entry-point into the test framework
//   golang has the testing framework built into it - you don't have to install a 3rd party one
// - the function name here has been prefixed with `Deprecated` so that this test no longer runs
//   as it is not recognised as a test any more

/*
func DeprecatedTestHelloWithSubtests(t *testing.T) {

	assertCorrect := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("bguiz")
		want := "Hello, bguiz"
		assertCorrect(t, got, want)
	})

	t.Run("say hello to default", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrect(t, got, want)
	})

}
*/

// NOTE
// - Using multiple `t.Run()` within within a `Test${something}` function is
//   like using a `describe` block to group multiple `it` tests in mocha JS
// - within `assertCorrect`, we use `t.Helper()` and this indicates that the
//   failing test is not this function within which `t.Errorf()` is called,
//   but rather the function which invoked it

func TestHelloWithLanguage(t *testing.T) {

	assertCorrect := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("bguiz", "")
		want := "Hello, bguiz"
		assertCorrect(t, got, want)
	})

	t.Run("say hello to someone in spanish", func(t *testing.T) {
		got := Hello("bguiz", "spanish")
		want := "Hola, bguiz"
		assertCorrect(t, got, want)
	})

	t.Run("say hello to someone in french", func(t *testing.T) {
		got := Hello("bguiz", "french")
		want := "Bonjour, bguiz"
		assertCorrect(t, got, want)
	})

	t.Run("say hello to default", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrect(t, got, want)
	})

	t.Run("say hello to default in spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola, World"
		assertCorrect(t, got, want)
	})

}

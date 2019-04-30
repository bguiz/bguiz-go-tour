package main

import (
	"fmt"
	"io"
	"net/http"
)

func GreetNoDi(name string) {
	fmt.Printf("Hello %s", name)
}

func GreetWithDi(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

func GreetHttpHandler(writer http.ResponseWriter, r *http.Request) {
	GreetWithDi(writer, "world")
}

func main() {
	http.ListenAndServe(":5555", http.HandlerFunc(GreetHttpHandler))
}

// NOTE
// - do `go run di.go` to start the webserver
// - this is a demonstration that by finding the right layer of abstraction
//   plus making use of DI, we can have the same function be used in multiple
//   scenarios. not just implementation and specification, but also for
//   multiple implmentations too (in this case a http server)

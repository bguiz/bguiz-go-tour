package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	var prefix string
	// if language == "spanish" {
	// 	prefix = spanishPrefix
	// } else {
	// 	prefix = englishPrefix
	// }
	switch language {
	case "spanish":
		prefix = spanishPrefix
	case "french":
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return prefix + name
}

// NOTE
// - `Hello` uses a capital H because that means it is public - can be seen by
//   other go files - think of it as exported.

func main() {
	fmt.Println(Hello("world", ""))
}

// NOTE
// - We separated `Hello` into its own function to aid with testability

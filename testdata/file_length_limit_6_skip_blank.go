package fixtures

import "fmt"

// Foo is a function.
func Foo(a, b int) {
	fmt.Println("Hello, world!")
}

// MATCH /file length is 7 lines, which exceeds the limit of 6/
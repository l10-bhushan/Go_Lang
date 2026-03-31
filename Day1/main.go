// Whenever we are creating a individual file we add package main
package main

// This package allows us to output onto our terminal
import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
	// Since both files come under same package, no need to import
	// We can directly use the function here
	// But to run this you need to create a go.mod file outside
	// and then run "go run ." command in Day1 folder"
	datatypes()
	conditionals()
	loops()
	arrays()
}

// 'go run main.go' allows us to run code directly
// 'go build main.go' builds a executable file
// ./main runs the executable

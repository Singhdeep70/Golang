package main

// Importing the "fmt" package to use the Println function for output
import "fmt"

func main() {
	// During compilation, Go infers the type as string and allocates memory on the stack
	// A string header is created with a pointer to the actual text "Happy birthday! You are now"
	messageStart := "Happy birthday! You are now"

	// Go infers 'age' as int (typically int32 or int64 depending on system)
	// Memory is reserved for the integer value 21
	age := 21

	// Another string variable is declared and memory is allocated similarly
	messageEnd := "years old!"

	// Behind the scenes, fmt.Println does the following:
	// 1. Converts non-string values (like 'age') to strings
	// 2. Joins all arguments with spaces
	// 3. Sends the final string to standard output (stdout)
	//    This means the text is passed to the OS, which sends it to the console or terminal
	fmt.Println(messageStart, age, messageEnd)

	// After execution, Go handles automatic memory cleanup of local variables (via stack unwinding)
}

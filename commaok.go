package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	// Create a new reader that reads from standard input (keyboard)
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user
	fmt.Println("Enter the rating: ")

	// Read the input until a newline character ('\n') is encountered
	input, _ := reader.ReadString('\n') // The walrus operator := declares and initializes 'input'

	// Print a thank you message along with the input received
	fmt.Println("Thanks for rating,", input)
}

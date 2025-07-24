package main

import "fmt"

// main is the entry point of the Go program
func main() {
	// Declare an integer variable initialized to its zero value (0)
	var zeroInt int // 0

	// Declare a float64 variable initialized to its zero value (0.00)
	var zeroFloat float64 // 0.00

	// Declare a boolean variable initialized to its zero value (false)
	var zeroBool bool // false

	// Declare a string variable initialized to its zero value (empty string "")
	var zeroString string // ""

	// Print all variables using appropriate format verbs
	// %v     - default format (used for int and bool)
	// %.2f   - float with 2 decimal places
	// %q     - quoted string
	fmt.Printf("%v %.2f %v %q\n", zeroInt, zeroFloat, zeroBool, zeroString)
}

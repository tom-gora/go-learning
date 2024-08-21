package lesson3

import "fmt"

// Run ...
func Run() {
	fmt.Print("Hello from lesson3")
}

// FloatToIntTruncation ...
func FloatToIntTruncation(f float64) int {
	return int(f)
}

// Sum ...
func Sum(a int, b int) int {
	return a + b
}

// Concat ...
func Concat(a string, b string) string {
	return a + " " + b
}

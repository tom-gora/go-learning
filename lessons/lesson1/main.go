// Package lesson1 Lesson 1 - Hello World
package lesson1

import "fmt"

// Run hello world
func Run() {
	fmt.Println("hello from lesson 1")
}

// DoSomethingElse is an another function
func DoSomethingElse() {
	n := 2
	for i := 0; i <= n; i++ {
		fmt.Printf("%v. One more line!\n", i)
	}
}

// Divide ...
func Divide(a int, b int) float64 {
	aFloat, bFloat := float64(a), float64(b)
	return aFloat / bFloat
}

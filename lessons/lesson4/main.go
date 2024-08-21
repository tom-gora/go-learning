// Package lesson4 contains functions for lesson 4
package lesson4

import (
	"fmt"
	"math/rand"
)

// Run ...
func Run() {
	fmt.Println("Hello from lessonm 4")
}

// ReturnRandomInt ...
func ReturnRandomInt(min, max int) int {
	return min + 1 + rand.Intn(max-min-1)
}

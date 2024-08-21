// Package lesson6 is a lesson 6 package
package lesson6

import "fmt"

// Run print welcome msg
func Run() {
	fmt.Printf("Hello from lesson 6\nWe are honing in structs now!\n")
}

// DoASlice slices off a part of a string. Written to learn about go's slicing syntax
func DoASlice(target string, direction string, charAmount int) (string, error) {
	if target == "" {
		return "", fmt.Errorf("string is empty")
	}
	if charAmount < 0 {
		return "", fmt.Errorf("cannot slice a negative amount of characters")
	}
	if charAmount >= len(target) {
		return "", fmt.Errorf("cannot slice a string that is shorter than the amount of characters to cut off")
	}
	if direction != "start" && direction != "end" {
		return "", fmt.Errorf("invalid direction")
	}
	// cut off charAmount from end of string but in fact this mean retain the amount of chars that come after :
	// and throw away the rest. That's why this num is calculated by taking entire lenght minus desired num to cut off
	if direction == "end" {
		return target[:len(target)-charAmount], nil
	}
	// cut off charAmount from start of string (imagine : stands for the string)
	if direction == "start" {
		return target[charAmount:], nil
	}
	return "", fmt.Errorf("internal error")
}

// NOTE: declare method on struct

// Cat - first declare some simple struct
type Cat struct {
	Name      string
	Age       int
	Lovliness float64
	HasEaten  bool
}

// Feed method declared on struct of type Cat
func (c *Cat) Feed() {
	c.HasEaten = true
	if c.Lovliness < 0.5 && c.HasEaten {
		c.Lovliness = 0.8
		fmt.Printf("You have fed %s and his loveliness increased to %.1f\n", c.Name, c.Lovliness)
		return
	}
	if c.Lovliness >= 0.5 && c.HasEaten {
		c.Lovliness = c.Lovliness + 0.2
		fmt.Printf("You have fed %s and his loveliness increased to %.1f\n", c.Name, c.Lovliness)
	}
}

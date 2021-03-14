// # Hands on Exercises - Level 2 #6
//
// Using iota, create 4 constants for the last 4 years. Print the constant values.
//
// [answer](https://play.golang.org/p/ffzD0s7w4m)

package main

import (
	"fmt"
)

// For future me:
// I tried to convert time.Year() into a constant, without success.
// Also, I can't see the reason of constants here. Well scoped vars make much more sense.
// Now, prove me wrong.

const (
	a = 2021 - iota
	b
	c
	d
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

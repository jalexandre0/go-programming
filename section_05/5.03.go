// # Hands on Exercises - Level 2 #3
//
// Create TYPED and UNTYPED constants. Print the values of the constants.
//
// [answer](https://play.golang.org/p/nZit_5l3fP)

package main

import (
	"fmt"
)

const (
	x     = 10 // untyped constant
	y int = 10 // typed constant
)

func main() {
	fmt.Printf("x: %v, y: %v", x, y)
}

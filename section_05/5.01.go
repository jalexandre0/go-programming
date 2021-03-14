// # Hands on Exercises - Level 2 #1
//
// Write a program that prints a number in decimal, binary, and hex.
//
// [answer](https://play.golang.org/p/ETreejfgse)

package main

import "fmt"

func main() {
	n := 10
	fmt.Printf("Decimal: %d\n", n)
	fmt.Printf("Binary: %b\n", n)
	fmt.Printf("Hex: %#x\n", n)
}

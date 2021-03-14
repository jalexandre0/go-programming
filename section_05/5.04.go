// # Hands on Exercises - Level 2 #4
//
// Write a program that
//   * assigns an int to a variable
//   * prints that int in decimal, binary, and hex
//   * shifts the bits of that int over 1 position to the left, and assigns that to a variable
//   * prints that variable in decimal, binary, and hex
//
//   [answer](https://play.golang.org/p/8BKIiPD4Ud)

package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

	y := x << 1

	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%#x\n", y)

}

// # Hands on Exercises - Level 3 #7
//
// Building on the previos exervice, create a program that uses else if and else.
//
// [answer](https://play.golang.org/p/6eIoEuhYmL)

package main

import (
	"fmt"
)

func main() {
	n := "James"

	if n == "Joe" {
		fmt.Printf("Hello %s \n", n)

	} else if n == "Jack" {
		fmt.Printf("Hello %s \n", n)
	} else {
		fmt.Printf("Who are you, %s ?  \n", n)
	}
}

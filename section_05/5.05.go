// # Hands on Exercises - Level 2 #5
//
// Create a variable of type string using a raw string literal, and then print it.
//
// [answer](https://play.golang.org/p/wj4t4vcnUa)
//

package main

import "fmt"

func main() {
	s := `This is a "Literal String": 
		  Literal strings are raw and respect everything into backticks, 		 
		  even returns 
		  										and tabs.`
	fmt.Printf("%s", s)
}

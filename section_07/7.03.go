// # Hands on Exercises - Level 3 #3
//
// Create a for loop using this syntax
//
//   * for condition {}
//
// Have it print out the years you have been alive.
//
// [answer](https://play.golang.org/p/tnyqBPJ-i5)

package main

import (
	"fmt"
	"time"
)

func main() {
	born_year := 1983
	year := time.Now().Year()

	for born_year <= year {
		fmt.Println(born_year)
		born_year++
	}
}

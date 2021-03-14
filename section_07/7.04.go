// # Hands on Exercises - Level 3 #4
//
// Create a for loop using this syntax
//
//   * for {}
//
// Have it print out hte years you have been alive.
//
// [answer](https://play.golang.org/p/ujil0hqEmK)

package main

import (
	"fmt"
	"time"
)

func main() {
	born_year := 1983
	year := time.Now().Year()

	for {
		if born_year > year {
			break
		}

		fmt.Println(born_year)
		born_year++
	}
}

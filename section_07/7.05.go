// # Hands on Exercises - Level 3 #5
//
// Print out the remainder modulus which is found for each number between 10 and 100 when it is divided by 4.
//
// [answer](https://play.golang.org/p/ohfJOW9euy)

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		m := i % 4
		fmt.Printf("%d, %d\n", i, m)
	}
}

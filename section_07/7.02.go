// # Hands on Exercises - Level 3 #2
//
// Print every rune code point of hte uppercase alphabet three times. Your output should look like this:
//
// 65
//   U+0041 'A'
//   U+0041 'A'
//   U+0041 'A'
//
// 66
//   U+0042 'B'
//   U+0042 'B'
//   U+0042 'B'
//
// ... through the rest of the characters.
//
// [answer](https://play.golang.org/p/-DRL4e60jD)

package main

import "fmt"

func main() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%+U, %q\n", i, i)
		}
		fmt.Println()
	}
}

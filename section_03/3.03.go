// Hands-on Exercise #3
//
// Using the code from the [previous exercise](https://play.golang.org/p/N3xbuFcxMy),
//
//   1. at the package level scope, assign the following values to the three variables
//     a. for `x` assign `42`
//     b. for `y` assign `"James Bond"`
//     c. for `z` assign `true`
//
//   2. in `func main`
//     a. use `fmt.Sprintf` to print all of the VALUES to one single string. ASSIGN the returned value of type string using the short declaration operator to a VARIABLE with the IDENTIFIER `s`
//
// [answer](https://play.golang.org/p/1HUBT9-WDo)

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%d \t  %s \t %t", x, y, z)
	fmt.Println(s)

}

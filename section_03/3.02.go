//  Hands-on Exercise #2
//
// 1. Use var to DECLARE three variables. The variables should have package level scope.
// Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the follow TYPE:
//   a. identifier `x` type int
//   b. identifier `y` type int
//   c. identifier `z` type int
//
// 2. in `func main`
//   a. print out the values for each identifier
//   b. the compiler assigned values to the variables. What are these values called?
//
// [answer](https://play.golang.org/p/N3xbuFcxMy)

package main

import "fmt"

var x int
var y int
var z int

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// in golang, they are called 'zero value'  or default zero value.

}

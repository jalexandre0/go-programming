// # Hands-on Exercise #5
//
// Building on the code from the [previous example](https://play.golang.org/p/IXcSrHDVu2)
//   1. at the package level scope, using the `var` keyword, create a VARIABLE with the IDENTIFIER `y`.
// The variable should be of the UNDERLYING TYPE of your custom type `x`.
//   2. in `func main`
//     a. this should already be done
//       - print out the value of the variable `x`
//       - print out the type of the variable `x`
//       - assign your own VALUE to the VARIABLE `x` using the `=` OPERATOR
//       - print out the value of the variable `x`
//     b. now do this
//       - now use CONVERSION to convert the TYPE of the VALUE stored in `x` to the UNDERLYING TYPE
//         1. then use the short declaration operator to ASSIGN that value to `y`
//
// [answer](https://play.golang.org/p/cRMA1HvSeg)

package main

import "fmt"

type my_type int

var x my_type
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	//convert x (my_type) to y (int)

	y := uint(x)
	fmt.Println(y)

}

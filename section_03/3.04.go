// # Hands-on Exercise #4
//
// FYI: Documentation and new terminology "[underlying type](https://golang.org/ref/spec#Types)"
//
// For this exercise
//
//   1. create your own type. Have the underlying type be an `int`.
//   2. create a VARIABLE of your new TYPE with the IDENTIFIER `x` using the `var` keyword
//   3. in `func main`
//     a. print out the value of the variable `x`
//     b. print out the type of the variable `x`
//     c. assign `42` to the variable `x` using the `=` OPERATOR
//     d. print out the value of the variable `x`
//
// [answer](https://play.golang.org/p/IXcSrHDVu2)

package main

import "fmt"

type my_type int

var x my_type

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}

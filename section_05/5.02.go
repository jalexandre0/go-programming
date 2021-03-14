// # Hands on Exercises - Level 2 #2
//
// Using the following operators, write expressesions and assign their values to variables:
//
//   * `==`
//   * `<=`
//   * `>=`
//   * `!=`
//   * `<`
//   * `>`
//
//   [answer](https://play.golang.org/p/dRF6_bk5kD)

package main

import "fmt"

func main() {
	a := (1 == 1)
	b := (1 <= 1)
	c := (2 >= 2)
	d := (3 != 4)
	e := (3 < 4)
	f := (4 > 3)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}

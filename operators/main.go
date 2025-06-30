// this script displays how precedence is used in operators
package main

import "fmt"

func main() {
	// precedence
	// multiplication
	a := 12.0 * 3.0 / 4.0
	b := (12.0 * 3.0) / 4.0
	c := 12.0 * (3.0 / 4.0)

	fmt.Println("a", a, "b", b, "c", c)

	// integer division
	unclear := 12 * (3 / 4)
	fmt.Println("unclear", unclear)

	// parenthesis
	f := 12.0 / 3.0 / 4.0
	fmt.Println("f", f)
	f = 12.0 / (3.0 / 4.0)
	fmt.Println("f", f)

	// addition/subtraction
	fmt.Println()
	x := 12 + 3 - 4
	y := (12 + 3) - 4
	z := 12 + (3 - 4)
	fmt.Println("x", x, "y", y, "z", z)

}

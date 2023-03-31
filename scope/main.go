package main

import "fmt"

var one = "One"

func main() {
	var somethingElse = "this is a block level variable"
	fmt.Println(somethingElse)
	myFunc()

}

func myFunc() {
	fmt.Println(one)
}

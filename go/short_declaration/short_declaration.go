package main

import "fmt"

var i = 10

func main() {

	//score := 0 //Don't use this
	var score int

	fmt.Println(score)

	width, height := 100, 50

	fmt.Println("width is :", width)
	fmt.Println("height is :", height)

	fmt.Println("package's i:", i)

	// package's i is being shadowed:
	i := 5
	// i above is a new variable.

	fmt.Println("main block:", i)

	// a new scope begins
	{
		// i is a new variable
		i := 6
		fmt.Println("inner block:", i)
	}
	// the scope ends

	// main's i
	fmt.Println("main's i:", i)
}

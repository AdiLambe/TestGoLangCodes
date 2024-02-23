package main

import "fmt"

func greet() func() {
	// variable is declared in outer function

	return func() {
		fmt.Println("Hi john")
	}

}

func main() {
	x := [3]string{"a", "b", "c"}
	y := x[0:1]
	fmt.Println(y)
}

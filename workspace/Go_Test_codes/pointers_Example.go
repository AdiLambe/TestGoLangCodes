/*
In Go, a pointer is a variable that stores the memory address of another variable.
Pointers allow us to directly work with memory addresses,
enabling us to access and modify the values stored in memory.

Memory address in Go:
-> When we create a variable, a memory address is allocated for that variable
to store its value.
-> we can access the memory address using the "& operator"

	Example : var num int = 5
			  fmt.Println("variable value", num)
			  fmt.println("memory address", &num)

-> we use pointer variables to store memory addresses.

	Example : var num := 5
			  var ptr *int = &num
	Here, "ptr" is a pointer variable that stores memory address of the "num" variable.
	"*int" indicates that "ptr" is of type int (it stores the memory address of an int variable)
	-> we can create pointer of other types as well (eg: *string, *float64 etc)

-> We can assign memory address of a variable to a pointer.

	var name = "john"
	var ptr *string
	ptr = &name
	fmt.Println("value of pointer is", ptr)
	fmt.Println("Address of variable is", &name)
*/
package main

import "fmt"

func main() {
	var name = "john"
	var ptr *string = &name
	fmt.Println(*ptr)
	//*ptr --> will give us the value present on that address
}

/*Here *ptr access the value stored in the memory address pointed by the pointer
(which is the value of "name")

point to remember : pointer in go allows us to work directly with memory, making them
powerful tools for efficient programming.*/

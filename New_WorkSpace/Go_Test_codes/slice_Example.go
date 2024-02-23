/*
In go, Slice is a dynamic and flexible data structure that represents
a segment of an array.
Unlike arrays, slice don't have a fixed size, and you caneasily add
or remove elements from them. below is the example of slice.
*/
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers :", numbers)
}

//we declare a slice named "numbers" containing integers.
//[] int -> specifies that the slice can only store integer values.
//output will be: "Numbers : [1 2 3 4 5]"

/*key points:
(1) Slices are created using the empty square brakets []
(2) If we provide a size inside the brakets, it becomes array
	Example: numbers := [5]int{1,2,3,4,5} -> created an Array
			 numbers := []int{1,2,3,4,5} -> created a slice
(3) you can also create a slice from an existing array by specifying
a range of indices.
Example: slicenumber := numbers[1:4]
-> it creates a slice including the elements at indices 1,2 and 3
	i.e {2,3,4}*/

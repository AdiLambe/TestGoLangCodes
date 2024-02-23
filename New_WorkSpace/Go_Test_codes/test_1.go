// /* Factorial */
// package main

// import "fmt"

// func factorial(n int) int {
// 	if n < 0 {
// 		fmt.Printf("factorial is undefined for negative number")
// 		return 0
// 	} else if n == 0 {
// 		return 1
// 	} else {
// 		return n * factorial(n-1)
// 	}
// }
// func main() {
// 	var num int
// 	fmt.Scan(&num)
// 	result := factorial(num)
// 	fmt.Println("output:", result)
// 	/*or we can simply write*/
// 	/*fmt.Println(factorial(5))*/
// }

//-----------------------------------------------------------------
/*Fibonacci series*/
// package main

// import "fmt"

// func fibonacci(n int) int {
// 	if n < 2 {
// 		return n
// 	} else {
// 		return fibonacci(n-1) + fibonacci(n-2)
// 	}
// }
// func main() {
// 	var num int
// 	fmt.Scan(&num)
// 	result := fibonacci(num)
// 	fmt.Println(result)
// 	// or we can simply  write
// 	// fmt.Println(fibonacci(7))
// }

//-----------------------------------------------------------------
/* Reverse a string */
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func reverseString(s string) string {
// 	// line 57 : converts input string s into a rune slice (since string in Go are immutable)
// 	/* A rune slice is a collection of runes.
// 	   when you convert a string to a rune slice, you get a new slice that contains the unicode code points of the string.
// 	   Example : we have string as "hello".
// 	   Converting it to rune slice gives us [104 101 108 108 111],
// 	   where each number represents the unicode code point of the correspomding character.
// 	*/
// 	runes := []rune(s)
// 	// It use two pointers ("i" and "j") to characters from both ends of the rune slice until they meet in middle.
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	// finally, it converts the reversed rune slice back to a string and returns it.
// 	return string(runes)
// }
// func main() {
// 	input := "hello"
// 	// we call the reverseString func with the input string and store the result in the "reversed" variable.
// 	reversed := reverseString(input)
// 	fmt.Println("Input:", input)
// 	fmt.Println("Output:", reversed)
// }

//-----------------------------------------------------------------
// /* Check if number is Prime
// Approach :
// 	-> Iterate over all positive integer less than the given number.
// 	-> Check if any of them divide the number evenly (i.e with no remainder)
// 	-> If none of them divide the number evenly, then the number is prime.
// */

// package main

// import "fmt"

// func isPrime(n int) bool {
// 	// It checks if the input number "n" is less than or equal to 1. Is so, return false (0 and 1 is not prime)
// 	if n <= 1 {
// 		return false
// 	}
// 	// For each value of "i", it checks if "n"is divisible evenly by "i". If so, it returns "false"
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	// If no divisor is found, it returns "true", indicating that "n" is prime.
// 	return true
// }

// func main() {
// 	input := 7
// 	// we call isPrime func with the input number and store the result in the "result" variable.
// 	result := isPrime(input)
// 	fmt.Printf("Is %d is a prime number? %v\n", input, result)
// }

//-----------------------------------------------------------------
/* Swap two numbers without using third variable  */
// package main

// import "fmt"

// func swapUsingArithematic(a, b int) (int, int) {
// 	a = a + b
// 	b = a - b
// 	a = a - b

// 	return a, b

// }

// func SwapUsingBitWiseOperator(a, b int) (int, int) {
// a = 5, b = 10
// 	a = a ^ b // initally, a contains the value a ^ b
// 	b = a ^ b // cancel out original value of "b"
//  a = a ^ b // cancel out original value of "a"

// 	return a, b
// }

// func SwapUsingMultiplicationDivision(a, b int) (int, int) {
// 	a = a * b
// 	b = a / b
// 	a = a / b

// 	return a, b
// }

// func main() {
// 	a, b := 5, 10
// 	fmt.Println("Before swapping:", a, b)
// 	a, b = SwapUsingMultiplicationDivision(a, b)
// 	fmt.Println("After swapping:", a, b)
// }

//----------------------------------------------------
//remember :
// func main() {
// 	a, b := 3, 4
// 	a, b = b, a
// 	fmt.Println("output: ", a, b)
// }

/*
	temp = a
	a = b
	b = temp
*/

// Swap two variables using third variable

// func swap(a *int, b *int) {
// 	var temp int

// 	temp = *a
// 	*a = *b
// 	*b = temp
// }

// func main() {
// 	var x, y int
// 	fmt.Println("Enter values for a and b:")
// 	fmt.Scan(&x, &y)
// 	swap(&x, &y)
// 	fmt.Println("output:", x, y)

// }

// -----------------------------------------------------------------
// /* Find max number in an array */
// package main

// import "fmt"

// func findMax(arr []int) int {
// 	max := arr[0]

// 	for _, num := range arr {
// 		if num > max {
// 			max = num
// 			fmt.Println("Max Number : ", max)
// 		}
// 	}
// 	return max
// }
// func main() {
// 	input := []int{3, 7, 1, 9, 4}
// 	result := findMax(input)

// 	fmt.Println("output:", result)
// }

//-----------------------------------------------------------

// package main

// import "fmt"

// func modifyString(s *string) {
// 	*s = "Modified"
// }
// func main() {
// 	str := "Original"
// 	strPtr := &str
// 	fmt.Println("Before modification:", str)
// 	modifyString(strPtr)
// 	fmt.Println("After modification:", str)
// }

//------------------------------

// package main

// import "fmt"

// func intersection(a, b []int) []int {

// 	c := make([]int, 0)
// 	for _, num1 := range a {
// 		for _, num2 := range b {
// 			if num1 == num2 {
// 				c = append(c, num1)
// 			}
// 		}
// 	}
// 	return c
// }

// func main() {
// 	a := []int{1, 2, 3, 4}
// 	b := []int{3, 4, 5, 6}

// 	result := intersection(a, b)
// 	fmt.Println("Output:", result)

// }

//--------------------------------------------

//Sort Array of Integers:
//Write a function to sort an array of integers in ascending order.

// package main

// import (
// 	"fmt"
// 	//"sort"
// )

// func sortIntegers(arr []int) {
// 	//sort.Ints(arr)
// 	n := len(arr)
// 	count := 0
// 	flag := 0
// 	for i := 0; i < n; i++ {
// 		flag++
// 		fmt.Println(" 'i' flag :", flag)
// 		for j := 0; j < n-i-1; j++ {
// 			count++
// 			fmt.Println("'j' count :", count)
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// }

// func main() {
// 	a := []int{5, 2, 9, 1, 4, 6}
// 	fmt.Println("Unsorted Array", a)
// 	sortIntegers(a)

// 	fmt.Println("Sorted Array", a)
// }

//--------------------------------------------

// input = 14578
// output = 25

// package main

// import "fmt"

// func sumofDigits(n int) int {
// 	var s int
// 	for n > 0 {
// 		s = s + n%10
// 		n = n / 10
// 	}
// 	return s
// }

// func main() {
// 	n := 14578
// 	result := sumofDigits(n)
// 	fmt.Println("Output", result)
// }

//------------------------------------------------

// -2, -3, 4, -1, -2, 1, 5, -3

// 6, -2, -3, -1, 4, -1, -2, 1, 5, -3, -4, 5, -1, 7, 8
package main

import (
	"fmt"
)

func isPrime(n int) bool {

	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	a := 7
	result := isPrime(a)
	fmt.Println("output :", result)
}

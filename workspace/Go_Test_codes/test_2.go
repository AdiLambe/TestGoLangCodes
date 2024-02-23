// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func reverseWords(s string) string {
// 	words := strings.Fields(s)                  // Split the sentence into words
// 	reversedWords := make([]string, len(words)) // Empty slice to reverse words

// 	// Reverse each word
// 	for i, word := range words {
// 		reversedWords[i] = reverseString(word)
// 	}

// 	// Join the reversed words back into a sentence
// 	return strings.Join(reversedWords, " ")
// }

// func reverseString(s string) string {
// 	runes := []rune(s) // Convert string to a slice of runes

// 	// Swap elements
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}

// 	return string(runes) // Convert back to a string
// }

// func main() {
// 	input := "Hello world This is golang"
// 	reversedSentence := reverseWords(input)

// 	fmt.Printf("Original sentence: %s\n", input)
// 	fmt.Printf("Reversed sentence: %s\n", reversedSentence)
// }

// ----------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func isPalindrome(s string) bool {
// 	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
// 	fmt.Println(len(s))
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	input := "Race Car"
// 	result := isPalindrome(input)
// 	fmt.Println("Output:", result)
// }

// // Hello World This is golang
// // olleH dlorW sihT si gnalog

//--------------------------------

// package main

// import (
// 	"fmt"
// )

// func maxSubArraySum(arr []int) int {
// 	n := len(arr)
// 	//fmt.Println(len(arr))
// 	maxEndingHere, maxSoFar := arr[0], arr[0]

// 	/*	•	This loop iterates through the array starting from the second element
// 			(since we already initialized maxEndingHere and maxSoFar with the first element).
// 		•	For each element, we calculate the new maxEndingHere by choosing
// 			the maximum value between the current element and the sum of the previous maxEndingHere and the current element.
// 		•	We also update maxSoFar by choosing the maximum value between the current maxSoFar and the updated maxEndingHere.
// 	*/
// 	for i := 1; i < n; i++ {
// 		maxEndingHere = max(arr[i], maxEndingHere+arr[i])
// 		maxSoFar = max(maxSoFar, maxEndingHere)
// 	}

// 	return maxSoFar
// }

// func main() {
// 	arr := []int{6, -2, -3, -1, 4, -1, -2, 1, 5, -3, -4, 5, -1, 7, 8}
// 	result := maxSubArraySum(arr)
// 	fmt.Printf("Maximum subarray sum: %d\n", result)
// }

//--------------------------------------------------------

// Kadane's algorithm

// Print the Maximum Subarray Sum: Given an array of integers,
// find the contiguous subarray with the maximum sum and print the subarray itself.
// Example: Input: arr = [-2, -3, 4, -1, -2, 1, 5, -3]
// Output: [4, -1, -2, 1, 5]

// package main

// import "fmt"

// func maxSubArraySum(arr []int) []int {
// 	n := len(arr)
// 	maxEndingHere, maxSoFar := arr[0], arr[0]
// 	startIndex, endIndex := 0, 0

// 	// this loop iterates through the array strting from second element
// 	for i := 1; i < n; i++ {
// 		if arr[i] > arr[i]+maxEndingHere {
// 			maxEndingHere = arr[i]
// 			startIndex = i
// 		} else {
// 			maxEndingHere += arr[i]
// 		}
// 		if maxEndingHere > maxSoFar {
// 			maxSoFar = maxEndingHere
// 			endIndex = i
// 		}
// 		fmt.Println(arr[startIndex : endIndex+1])
// 	}
// 	return arr[startIndex : endIndex+1]
// }

// func main() {
// 	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
// 	result := maxSubArraySum(arr)
// 	fmt.Println("output: ", result)
// }

//-----------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	//H"strings"
// )

// func modifyString(s *string) {
// 	*s = "Modified String"
// }

// func main() {
// 	var input string
// 	fmt.Println("Enter Input:")
// 	//input = strings.ReplaceAll("Enter input string :", " ", " ")
// 	fmt.Scanln(&input)
// 	modifyString(&input)
// 	fmt.Println("Output:", input)
// }

//------------------------------------------

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func modifyString(s *string) {
// 	*s = "Modified String"
// }

// func main() {
// 	var input string
// 	fmt.Println("Enter Input:")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ = reader.ReadString('\n')
// 	input = strings.TrimSpace(input)
// 	modifyString(&input)
// 	fmt.Println("Output:", input)
// }

// --------------------------------------------
package main

import "fmt"

func sortIntegers(arr []int) {
	n := len(arr)
	count := 0
	flag := 0

	for i := 0; i < n; i++ {
		flag++

		for j := 0; j < n-i-1; j++ {
			count++
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	a := []int{5, 2, 9, 1, 4, 6}

	sortIntegers(a)
	fmt.Println("sorted array :", a)
}

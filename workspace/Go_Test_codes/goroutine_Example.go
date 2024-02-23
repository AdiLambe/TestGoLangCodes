/*In Go, a goroutine is a lightweight thread for execution that runs "concurrency" with other goroutines.
Goroutines allow you to write "concurrent code" easily, making your program more effective and responsive.

	What is goroutine ?
	A goroutine is a fundamental feature of golang.
	It represents an indepedently executing function or method.
	Goroutines are lightweight compared to traditional operating system threads.

	Why use goroutine ?
	Concurrency : Goroutines enable concurrent execution of tasks. Multiple goroutines can run simultaneously,
	performing different tasks.
	Efficiency : Creating and managing goroutines is cheap in terms of memory and resources.
	Scalability : You can create thousands or even millions of goroutines without much overload.
	Responsive Programs : Goroutine allow non-blocking execution, making programs more responsive

	Creating a Goroutine :
		To create a goroutine, use the "go" keyword followed by a function call.
*/

package main

import (
	"fmt"
	"time"
)

// Defines a func name "printNumbers"
func printNumbers() {
	// It prints number from 1 to 5
	// Adds a delay of milliseconds between each print using time.sleep
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 800)
	}
}
func main() {

	// create goroutine
	go printNumbers()

	// Main goroutine continues executing the loops that print "Main : 1", "Main : 2" etc
	// It also adds a delay of 200 milliseconds between each print.
	for i := 1; i <= 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(time.Millisecond * 400)
	}
}

/*
Goroutine Lifecycle :
	Goroutines run concurrently with the main program.
	When the main Program exits, all remaining goroutines are terminated.
Communication between goroutine :
	Goroutines can communicate using channels (a built-in synchronization mechanism)
	Channels allow safe data sharing between goroutines.

In summary, goroutines are a powerful way to achieve concurrency in go, making it easy to write efficient and responsive programs.
*/

/*
In Go, channel are a powerful mechanism for communication and synchronization between
goroutines (concurrency executing function). They allow goroutines to send and receive data safely.

channel:

	A channel is typed conduit through which you can send and receive value between goroutines.
	Channels ensure safe communication by synchronizing access to share data.
	You can think of a channel as a pipe connecting goroutines, allowing them to echange info.
	To create a channel, use the "make" function with the "chan" keyword and specify the type
	data it will carry.
*/
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 42
// 	}()

// 	value := <-ch
// 	fmt.Println("Rceived value from channel", value)
// }

/*Buffered channel :
By default, channels are "unbuffered", meaning they only accept sends if there's a corresponsding
receiver ready to receive the value.
Buffered channels allow you to send a limited number of values without an immediate receiver.
You specify the buffer capacity when creating a channel using "make(chan type, capacity)" */

package main

import (
	"fmt"
)

func main() {
	// Create an
	ch := make(chan int, 2)

	ch <- 10
	ch <- 36
	//ch <- 56

	fmt.Println("Received from channel 1:", <-ch)
	fmt.Println("Received from channel 2:", <-ch)
	//fmt.Println("Received from channel 3", <-ch)
}

/*Use Cases:
unbuffered channels are useful for precise synchronization between goroutines.
Buffered channels are handy when you want to decouple senders and receivers temporarily */

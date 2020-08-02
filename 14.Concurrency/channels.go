package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receives from c

	fmt.Println(x, y, x+y) // x+y will executed as soon as x and y are completed

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	d := make(chan int, 10)
	go fibonacci(cap(d), d)
	for i := range c {
		fmt.Println(i)
	}
}

/*

Channels are a typed conduit through which you can send and receive values with the channel operator, <-

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit
locks or condition variables.

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed
by assigning a second parameter to the receive expression.

Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

*/
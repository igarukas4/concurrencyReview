package main

import (
	"fmt"
	"time"
)

func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int, 5) // Buffered channel with a size of 5

	go produce(ch)
	go consume(ch)
	
	time.Sleep(1 * time.Second)
}

/*
4.2. Behavior Difference Between Buffered and Unbuffered Channels:
Unbuffered Channel: The sending goroutine blocks until the receiving goroutine receives the value. 
This means the sending and receiving happen simultaneously.
Buffered Channel: The sending goroutine does not block until the buffer is full. 
It allows the sending goroutine to continue execution until the buffer capacity is reached. 
The receiving goroutine can receive the values at its own pace. 
This can improve performance by decoupling the sending and receiving operations but requires managing the buffer size properly to avoid overfilling.
*/
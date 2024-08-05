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
	ch := make(chan int)

	go produce(ch)
	go consume(ch)
	
	// Sleep to ensure goroutines finish before the main function exits
	time.Sleep(1 * time.Second)
}
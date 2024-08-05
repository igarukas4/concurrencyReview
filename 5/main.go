package main

import (
	"fmt"
)

func produceNumbers(evenCh, oddCh chan int) {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go produceNumbers(evenCh, oddCh)

	for {
		select {
		case num, ok := <-evenCh:
			if ok {
				fmt.Printf("Received an even number: %d\n", num)
			}
		case num, ok := <-oddCh:
			if ok {
				fmt.Printf("Received an odd number: %d\n", num)
			}
		}
		if len(evenCh) == 0 && len(oddCh) == 0 {
			break
		}
	}
}

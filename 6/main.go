package main

import (
	"fmt"
)

func produceNumbers(evenCh, oddCh, errCh chan int) {
	for i := 1; i <= 25; i++ {
		if i > 20 {
			errCh <- i
		} else if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
	close(errCh)
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	errCh := make(chan int)

	go produceNumbers(evenCh, oddCh, errCh)

	evenChOpen, oddChOpen, errChOpen := true, true, true
	for evenChOpen || oddChOpen || errChOpen {
		select {
		case num, ok := <-evenCh:
			if ok {
				fmt.Printf("Received an even number: %d\n", num)
			} else {
				evenChOpen = false
			}
		case num, ok := <-oddCh:
			if ok {
				fmt.Printf("Received an odd number: %d\n", num)
			} else {
				oddChOpen = false
			}
		case err, ok := <-errCh:
			if ok {
				fmt.Printf("Error: number %d is greater than 20\n", err)
			} else {
				errChOpen = false
			}
		}
	}
}

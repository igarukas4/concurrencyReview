package main

import (
	"fmt"
)

func produceNumbers(evenCh, oddCh, errCh chan int) {
	for i := 1; i <= 22; i++ {
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
		case err, ok := <-errCh:
			if ok {
				fmt.Printf("Error: number %d is greater than 20\n", err)
			}
		}
		if len(evenCh) == 0 && len(oddCh) == 0 && len(errCh) == 0 {
			break
		}
	}
}

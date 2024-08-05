package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'j'; i++ {
		fmt.Println(string(i))
	}
}

func main() {
	go printNumbers()
	go printLetters()
	
	// Sleep to ensure goroutines finish before the main function exits
	time.Sleep(1 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 19; i++ {
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage("Goroutine")  // runs concurrently
	printMessage("Main function") // runs in main thread
	time.Sleep(500 * time.Millisecond)
}

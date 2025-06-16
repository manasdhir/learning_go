package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go printNumbers()

	for i := 1; i <= 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(300 * time.Millisecond)
	}

	//time.Sleep(1 * time.Second)
}

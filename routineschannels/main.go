package main

import (
	"fmt"
	"sync"
	"time"
)

func customer(id int, orders chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	order := fmt.Sprintln("Order from customer id", id)
	fmt.Println(order)
	orders <- order
}
func chef(id int, orders <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orders {
		fmt.Println("chef", id, "received", order)
		time.Sleep(time.Second)
		fmt.Println("chef", id, "prepared food", order)
	}
}

func main() {
	var chefwg sync.WaitGroup
	var customerwg sync.WaitGroup

	orders := make(chan string, 10)
	for i := 1; i <= 3; i++ {
		chefwg.Add(1)
		go chef(i, orders, &chefwg)
	}
	for i := 0; i <= 20; i++ {
		customerwg.Add(1)
		go customer(i, orders, &customerwg)
	}
	customerwg.Wait()
	close(orders)
	chefwg.Wait()
}

package main

import "fmt"

func main() {
	ic := 0
	for i := 0; i < 10; i++ {
		ic += i
	}
	fmt.Println(ic)

	var intcount, i uint8 = 0, 0

	for i < 10 {
		intcount += i
		i++
	}
	fmt.Println(intcount) // output 45
	for {
		fmt.Println("ASDAS") // infinite loop
	}
}

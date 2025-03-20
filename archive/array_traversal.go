package main

import "fmt"

func main() {

	arr1 := [200]int{1, 2, 3, 4, 5, 4, 6, 6, 7, 78, 8, 89, 9, 9, 0, 123, 4, 5, 5, 6}
	for i, num := range arr1 {
		fmt.Println(i, num)
	}

}

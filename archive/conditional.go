package main

import "fmt"

func main() {
	marks1 := 75
	if marks1 >= 75 {
		fmt.Println("Distinction")
	} else if marks1 < 75 && marks1 > 60 {
		fmt.Println("First Div")
	} else {
		fmt.Println("Others")
	}
	if marks := 69; marks >= 75 {
		fmt.Println("Distinction")
	} else if marks < 75 && marks > 60 {
		fmt.Println("First Div")
	} else {
		fmt.Println("Others")
	}
	day := 'A'
	switch day {
	case 65:
		fmt.Println("A Found")
		fallthrough
	case 66:
		fmt.Println("B Found")
	default:
		fmt.Println("Other")
	}
}

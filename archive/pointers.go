package main

import "fmt"

func double(ptrArr []*int) {
	for i := 0; i < len(ptrArr); i++ {
		*ptrArr[i] *= 2
	}
}
func main() {
	firstVar, secondVar, thirdVar := 1, 2, 3
	ptrArr := []*int{&firstVar, &secondVar, &thirdVar} // Array of pointers
	fmt.Println("Value of Each variable through Array of Pointers:")
	for i := 0; i < len(ptrArr); i++ {
		fmt.Println(*ptrArr[i]) // Dereferencing pointer
	}
	fmt.Println("Value of Each variable through Array of Pointers After Doubling:")
	double(ptrArr)
	for i := 0; i < len(ptrArr); i++ {
		fmt.Println(*ptrArr[i]) // Dereferencing pointer
	}
}

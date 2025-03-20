package main

import "fmt"

func main() {
	fmt.Println(intdiv(6, 4))
	fmt.Println((intdivwithrem(6, 4)))
	fmt.Println(intdivwithremnamed(6, 4))
	fmt.Println(summation(1, 2, 3, 4, 5, 6))
	fmt.Println(applyopr(4, 5, add))
	a := multiplier(5)
	fmt.Println(a((2)))
	b := multiplier(3)
	fmt.Println(b(2))
}

func intdiv(num, dem int) int {
	result := num / dem
	return result
}
func intdivwithrem(num, dem int) (int, int) {
	result := num / dem
	rem := num % dem
	return result, rem
}
func intdivwithremnamed(num, dem int) (quotient, remainder int) {
	quotient = num / dem
	remainder = num % dem
	return
}

func summation(num ...int) (total int) {
	for _, num := range num {
		total += num
	}
	return
}

func applyopr(num1, num2 int, op func(int, int) int) int {
	return op(num1, num2)
}

func add(num1, num2 int) int {
	return num1 + num2
}

func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

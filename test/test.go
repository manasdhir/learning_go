package main

import "fmt"

func main() {
	var ptr *int
	if ptr == nil {
		fmt.Println("Pointer is nil")
	}
	a, b := 1, 2
	fmt.Println(add(a, b))
	fmt.Println(div2(a, b))
	fmt.Println(div(a, b))
	fmt.Println(applyopr(a, b, div2))
	c := multiplier(4)
	fmt.Println(c(b))
	var stud1 = stud{"ssdfsd", 12}
	stud1.pr()
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
	var map1 = make(map[string]int)
	map1["asdasda"] = 123
	map1["sdfsdfs"] = 32131
	fmt.Println(map1)
	fmt.Println(map1["asdasda"])
	delete(map1, "asdasda")
	fmt.Println(map1)
	fmt.Println(summation(1, 2))
	fmt.Println(sumofseries(8, square))
	fmt.Println(sumofseries(5, cube))
	fmt.Println(sumofseries(2, multiplier(5)))
	var x, y, z *int = new(int), new(int), new(int)
	t := 5
	u := 6
	x = &t
	y = &u
	z = &t
	arr := [...]*int{x, y, z}
	fmt.Println(arr)
	arrptr := &arr
	fmt.Println(arrptr)
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr1ptr := &arr1
	arr1ptr[0] = 1214
	fmt.Println("array copy", arr2)
	fmt.Println("value from arrptr", *arr1ptr)
	fmt.Println("original array", arr1)

}

func add(num, dem int) int {
	return num + dem
}
func div(num, dem int) (int, int) {
	q := num / dem
	r := num % dem
	return q, r
}
func div2(num, dem int) (q, r int) {
	q = num / dem
	r = num % dem
	return
}
func applyopr(num int, dem int, a func(int, int) (int, int)) (int, int) {
	return a(num, dem)
}
func multiplier(num int) func(int) int {
	return func(num2 int) int {
		return num2 * num
	}

}

type stud struct {
	name   string
	rollno int
}

func (s stud) pr() {
	fmt.Println(s.name, "has roll no", s.rollno)

}

func double(ptrArr []*int) {
	for i := 0; i < len(ptrArr); i++ {
		*ptrArr[i] *= 2
	}
}
func summation(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func sumofseries(len int, op func(int) int) int {
	sum := 0
	for i := 0; i <= len; i++ {
		val := op(i)
		sum += val
	}
	return sum
}
func square(num int) int {
	return num * num
}
func cube(num int) int {
	return num * num * num
}

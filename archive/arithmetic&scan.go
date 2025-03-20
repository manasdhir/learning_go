package main

import "fmt"

func main() {
	var num1 int = 6
	var num2 int = 7
	fmt.Println(num1 / num2) // output will be 0 in integer type
	var f1 float32 = 6.943
	//fmt.Println((num1/f1)) // will result in error
	f2 := int(f1)
	fmt.Println((f2))        // output will be 6
	fmt.Println((num1 / f2)) //output will be 1
	var intvar16bit int16 = 4
	var intvar32bit int32 = int32(intvar16bit)
	fmt.Println(intvar32bit)
	var arrayElementNotInit = [3]int{} //Not Initialized  & type declaration and size
	var stringVar1, stringVar2 string
	fmt.Scanln(&stringVar1, &stringVar2) // Input from user upto space by space only and will exit on new line
	fmt.Println(stringVar1, " ", stringVar2)
	var name string
	var age int16
	fmt.Scanf("%s %d", &name, &age) // Works like Scanln but have format specifiers also.
	fmt.Println(name, age)

	fmt.Scan(&arrayElementNotInit[0], &arrayElementNotInit[1], &arrayElementNotInit[2]) // Input from user upto space by space or in next line separated also.
	fmt.Println(arrayElementNotInit)
}

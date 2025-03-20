package main

import "fmt"

func main() {
	var inum uint8
	fmt.Println(inum)
	var b byte = 'A'
	fmt.Println(b)
	var c rune = '😊'
	fmt.Println(c)
	var d bool
	fmt.Println(d)
	var e string = "asdas as"
	fmt.Println(e)
	var f = 6
	var j = 7
	fmt.Println(f + j)
	K := "wSADASDAS"
	fmt.Println(K)
	in1, in2, in3 := 1, 2, 3
	fmt.Print(in1, in2, in3)
	fmt.Print("\n")
	intVar, stringVar, boolVar := 1, "Hello", true
	fmt.Println(intVar, stringVar, boolVar)
	var (
		intVarBlock          int
		intVarBlockAss       uint = 9
		stringVarBlockAssInf      = "HelloInferrredBlock"
	)
	fmt.Println(intVarBlock, intVarBlockAss, stringVarBlockAssInf)
	const pi = 3.14
	fmt.Printf("%v %#v %T \n", intVarBlock, stringVarBlockAssInf, intVarBlockAss)

}

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	slice = append(slice, 5)
	fmt.Println(slice)
	stud1 := student{"asdasa", 1}
	fmt.Println(stud1.studname, stud1.studrollno)
	stud1.pr()
	stud2 := student{"fdsfsdf", 2}
	stud2.pr()
}

type student struct {
	studname   string
	studrollno int
}

func (s student) pr() {
	fmt.Println("student name:", s.studname, "student roll no.:", s.studrollno)
}

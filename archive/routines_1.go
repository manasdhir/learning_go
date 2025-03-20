package main

import (
	"fmt"
	"time"
)

type Student struct {
	studName       string
	studRollNumber int32
}

func (s Student) print() {
	time.Sleep(time.Second)
	fmt.Println("Student Name:", s.studName, "Roll:", s.studRollNumber)
}
func main() {
	stud1 := Student{"Rajesh", 1}
	go stud1.print()
	stud3 := Student{"Rajesh2", 3}
	go stud3.print()
	stud2 := Student{"Ramesh", 2}
	stud2.print()
	time.Sleep(time.Second)
}

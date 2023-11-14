package main

import "fmt"

type Student struct {
	rno, marks int
	sname      string
}

func (s *Student) Accept() {
	fmt.Println("ENTER ROLL NO:::")
	fmt.Scanln(&s.rno)
	fmt.Println("STUDENT NAME:::")
	fmt.Scanln(&s.sname)
	fmt.Println("MARKS:::")
	fmt.Scanln(&s.marks)
}

func (s *Student) Show() {
	fmt.Println("ROLL NO:::", s.rno)
	fmt.Println("STUDENT NAME:::", s.sname)
	fmt.Println("MARKS:::", s.marks)
}

func main() {
	var obj Student
	obj.Accept()
	obj.Show()
}

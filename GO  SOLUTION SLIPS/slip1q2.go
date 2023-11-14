package main

import "fmt"

type Student struct {
	rno                    int
	sname                  string
	m1, m2, m3, total, avg int
}

func (s Student) Display() {
	fmt.Println("ROLL NO:::", s.rno)
	fmt.Println("STUDENT NAME:::", s.sname)
	fmt.Println("MARKS 1:::", s.m1)
	fmt.Println("MARKS 2:::", s.m2)
	fmt.Println("MARKS 3:::", s.m3)
	fmt.Println("TOTAL:::", s.total)
	fmt.Println("AVERAGE:::", s.avg)
}
func main() {
	var sname string
	var rno, m1, m2, m3, total, avg int
	fmt.Println("ENTER ROLL NO")
	fmt.Scanln(&rno)
	fmt.Println("ENTER STUDENT NAME")
	fmt.Scanln(&sname)
	fmt.Println("ENTER MARKS OF SUBJECT 1")
	fmt.Scanln(&m1)
	fmt.Println("ENTER MARKS OF SUBJECT 2")
	fmt.Scanln(&m2)
	fmt.Println("ENTER MARKS OF SUBJECT 3")
	fmt.Scanln(&m3)
	total = m1 + m2 + m3
	avg = total / 3
	obj := Student{rno, sname, m1, m2, m3, total, avg}
	obj.Display()
}

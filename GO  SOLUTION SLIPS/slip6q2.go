package main

import "fmt"

type Student struct {
	rno        int
	name       string
	per, total int
	m1, m2, m3 int
}

func (s *Student) Accept() {
	fmt.Println("ENTER ROLL NO:::")
	fmt.Scanln(&s.rno)
	fmt.Println("ENTER NAME:::")
	fmt.Scanln(&s.name)
	fmt.Println("ENTER MARKS 1:::")
	fmt.Scanln(&s.m1)
	fmt.Println("ENTER MARKS 2:::")
	fmt.Scanln(&s.m2)
	fmt.Println("ENTER MARKS 3:::")
	fmt.Scanln(&s.m3)
	s.total = s.m1 + s.m2 + s.m3
	s.per = s.total / 3
}

func (s *Student) Display() {
	fmt.Println("ROLL NO:::", s.rno)
	fmt.Println("NAME:::", s.name)
	fmt.Println("MARKS 1:::", s.m1)
	fmt.Println("MARKS 2:::", s.m2)
	fmt.Println("MARKS 3:::", s.m3)
	fmt.Println("TOTAL MARKS:::", s.total)
	fmt.Println("PERCENTAGE:::", s.per)
}
func main() {
	var n, i, j int
	var temp Student
	obj := make([]Student, 100)
	fmt.Println("ENTER THE NO OF DETAILS?")
	fmt.Scanln(&n)
	for i = 0; i < n; i++ {
		obj[i].Accept()
	}
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if obj[i].per < obj[j].per {
				temp = obj[i]
				obj[i] = obj[j]
				obj[j] = temp
			}
		}
	}
	fmt.Println("DETAILS IN DESCENDING ORDER OF PERCENTAGE:::")
	for i = 0; i < n; i++ {
		obj[i].Display()
	}
}

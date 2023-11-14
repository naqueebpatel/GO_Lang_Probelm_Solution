package main

import "fmt"

type Employee struct {
	eno, sal int
	ename    string
}

func (e Employee) Accept() {
	fmt.Println("ENTER EMPLOYEE NO:::")
	fmt.Scanln(&e.eno)
	fmt.Println("ENTER EMPLOYEE NAME:::")
	fmt.Scanln(&e.ename)
	fmt.Println("ENTER EMPLOYEE SALARY:::")
	fmt.Scanln(&e.sal)
}

func (e Employee) Display() {
	fmt.Println("EMPLOYEE NO:::", e.eno)
	fmt.Println("EMPLOYEE NAME:::", e.ename)
	fmt.Println("EMPLOYEE SALARY:::", e.sal)
}
func main() {
	var n, i int
	obj := make([]Employee, 100)
	fmt.Println("ENTER NO. OF DETAILS")
	fmt.Scanln(&n)
	for i = 0; i < n; i++ {
		obj[i].Accept()
	}
	max := 0
	x := 0
	for i = 0; i < n; i++ {
		obj[i].Display()
		if obj[i].sal > max {
			max = obj[i].sal
			x = i
		}
	}
	fmt.Println("EMPLOYEE WITH MAXIMUM SALARY IS:::", obj[x].ename)
}

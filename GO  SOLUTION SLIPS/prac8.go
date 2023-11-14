package main

import "fmt"

type Employee struct {
	id   int
	name string
	sal  float64
}

func (e *Employee) Accept() {
	fmt.Println("ENTER ID")
	fmt.Scanln(&e.id)
	fmt.Println("ENTER NAME")
	fmt.Scanln(&e.name)
	fmt.Println("ENTER SAL")
	fmt.Scanln(&e.sal)
}
func (e *Employee) Display() {
	fmt.Println("ID", e.id)
	fmt.Println("NAME", e.name)
	fmt.Println("SALARY", e.sal)
}

func main() {
	obj := make([]Employee, 100)
	var size int
	var upd string
	fmt.Println("ENTER NO OF DETAILS?")
	fmt.Scanln(&size)
	for i := 0; i < size; i++ {
		obj[i].Accept()
	}
	for i := 0; i < size; i++ {
		obj[i].Display()
	}
	fmt.Println("ENTER NAME TO UPDATE SALARY")
	fmt.Scanln(&upd)
	for i := 0; i < size; i++ {
		if obj[i].name == upd {
			obj[i].sal = obj[i].sal - (obj[i].sal * 0.1)
		}
	}
	fmt.Println("UPDATED SALARY")
	for i := 0; i < size; i++ {
		obj[i].Display()
	}
}

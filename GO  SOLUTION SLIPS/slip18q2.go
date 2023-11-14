package main

import "fmt"

func main() {
	var s1, s2 string
	var p1, p2 *string
	*p1 = s1
	*p2 = s2
	fmt.Println("ENTER STRING 1")
	fmt.Scanln(&s1)
	fmt.Println("ENTER STRING 2")
	fmt.Scanln(&s2)
	fmt.Scanln("CONCATENATION:::", *p1+*p2)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var mystr string
	var alpha string
	var cnt int
	fmt.Println("Enter String::")
	fmt.Scanln(&mystr)
	fmt.Println("Enter Alphabet To Find")
	fmt.Scanln(&alpha)
	cnt = strings.Count(mystr, alpha)
	fmt.Println("TOTAL COUNT OF ", alpha, " IS ", cnt)
}

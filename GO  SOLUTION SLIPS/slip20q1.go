package main

import "fmt"

func main() {
	var num, i, p int
	fmt.Println("ENTER NUMBER")
	fmt.Scanln(&num)
	fmt.Println("ENTER POWER TO RAISE")
	fmt.Scanln(&p)
	pow := 1
	for i = 1; i <= p; i++ {
		pow = pow * num
	}
	fmt.Println(num, " RAISED TO THE POWER ", p, " IS ", pow)
}

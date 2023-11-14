package main

import "fmt"

func CalcSq(num int, sq chan int) {
	sum := 0
	rem := 0
	for num > 0 {
		rem = num % 10
		sum = sum + rem*rem
		num = num / 10
	}
	sq <- sum
}
func CalcCube(num int, cb chan int) {
	sum := 0
	rem := 0
	for num > 0 {
		rem = num % 10
		sum = sum + rem*rem*rem
		num = num / 10
	}
	cb <- sum
}
func main() {
	cb := make(chan int)
	sq := make(chan int)
	go CalcSq(105, sq)
	go CalcCube(45, cb)
	sq1 := <-sq
	cb1 := <-cb
	fmt.Println("SUM:::", sq1+cb1)
}

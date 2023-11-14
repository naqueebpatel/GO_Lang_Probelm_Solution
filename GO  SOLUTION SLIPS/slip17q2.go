package main

import "fmt"

func SquareCalc(num int, sq chan int) {
	sum := 0
	rem := 0
	for num > 0 {
		rem = num % 10
		sum = sum + rem*rem
		num = num / 10
	}
	sq <- sum
}
func CubeCalc(num int, cb chan int) {
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
	var snum, cnum int
	sq := make(chan int)
	cb := make(chan int)
	fmt.Println("ENTER NUMBER FOR SQUARE")
	fmt.Scanln(&snum)
	fmt.Println("ENTER NUMBER FOR CUBE")
	fmt.Scanln(&cnum)
	go SquareCalc(snum, sq)
	go CubeCalc(cnum, cb)
	Sqsum := <-sq
	Cbsum := <-cb
	fmt.Println("FINAL SUMMATION:::", Sqsum+Cbsum)
}

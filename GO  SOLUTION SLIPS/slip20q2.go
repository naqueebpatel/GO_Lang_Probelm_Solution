package main

import (
	"fmt"
	"time"
)

func print() {
	for i := 0; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}
func main() {
	go print()
	time.Sleep(1300 * time.Millisecond)
	fmt.Println("MAIN TERMINATED")
}

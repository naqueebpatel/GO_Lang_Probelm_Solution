package main

import (
	"fmt"
	"time"
)

func f() {
	for i := 0; i <= 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
	}
}
func main() {
	go f()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("MAIN TERMINATED")
}

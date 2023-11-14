package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("MAIN STARTED")
	func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("INSIDE ANONYMOUS")
	}()
	fmt.Println("MAIN TERMINATED")

}

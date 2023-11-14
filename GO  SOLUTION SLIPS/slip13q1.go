package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("MAIN STARTED")
	func() {
		time.Sleep(1 * time.Second)
		fmt.Println("INSIDE ANONYMOUS ROUTINE")
	}()
	fmt.Println("MAIN TERMINATED")
}

package main

import "fmt"

func main() {
	var arr [10]int
	var size, i, j int
	fmt.Println("ENTER SIZE OF ARRAY:::")
	fmt.Scanln(&size)
	for i = 0; i < size; i++ {
		fmt.Println("ENTER ARRAY ELEMENT")
		fmt.Scanln(&arr[i])
	}
	for i = 0; i < size; i++ {
		for j = i + 1; j < size; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println("ARRAY ELEMENT IN ASCENDING ORDER:::")
	for i = 0; i < size; i++ {
		fmt.Println(arr[i])
	}

}

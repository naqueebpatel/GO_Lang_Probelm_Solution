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
	fmt.Println("SMALLEST ELEMENT IN ARRAY IS:::", arr[0])
	fmt.Println("LARGEST ELEMENT IN ARRAY IS:::", arr[size-1])

}

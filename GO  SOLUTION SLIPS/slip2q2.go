package main

import "fmt"

func main() {

	var matrixA [10][10]int
	var matrixB [10][10]int
	var matrixC [10][10]int
	var row, column, i, j int
	fmt.Println("ENTER THE NO OF ROWS FOR MATRIX")
	fmt.Scanln(&row)
	fmt.Println("ENTER THE NO OF COLUMNS FOR MATRIX")
	fmt.Scanln(&column)

	fmt.Println("ENTER ELEMENTS FOR MATRIX ONE")
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			fmt.Scanln(&matrixA[i][j])
		}
	}

	fmt.Println("ENTER ELEMENTS FOR MATRIX TWO")
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			fmt.Scanln(&matrixB[i][j])
		}
	}

	fmt.Println("MATRIX ONE:")
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			fmt.Print(matrixA[i][j], "\t")
		}
		fmt.Println()
	}

	fmt.Println("MATRIX TWO:")
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			fmt.Print(matrixB[i][j], "\t")
		}
		fmt.Println()
	}
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			matrixC[i][j] = matrixA[i][j] * matrixB[i][j]
		}
	}

	fmt.Println("MULTIPLIED MATRIX:")
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			fmt.Print(matrixC[i][j], "\t")
		}
		fmt.Println()
	}

}

package main

import (
	"fmt"
	"testing"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func Benchmarkfibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(10)
	}
}
func main() {
	fmt.Println(fibonacci(10))
	testing.Benchmark(Benchmarkfibonacci)
}

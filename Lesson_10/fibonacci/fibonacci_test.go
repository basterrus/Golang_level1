package fibonacci

import (
	"fmt"
	"testing"
)

//Реализация теста Benchmark
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Fibonacci(25)
	}
}

func BenchmarkFibonacciOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciOptimized(25, map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 3})
	}
}

//Реализация теста Example
func ExampleFibonacci() {
	fibonacci, err := Fibonacci(10)
	if err != nil {
		return
	}
	fmt.Println(fibonacci)

	//Output:
	//55
}

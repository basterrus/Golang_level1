package fibonacci

import "errors"

var ErrPositiveValue = errors.New("допускаются только положительные числа")

//Fibonacci Function for calculation fibonacci numbers
func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, ErrPositiveValue
	}
	if n <= 1 {
		return n, nil
	}

	val1, err := Fibonacci(n - 1)
	if err != nil {
		return 0, err
	}
	val2, err := Fibonacci(n - 2)
	if err != nil {
		return 0, err
	}

	return val1 + val2, nil
}

//FibonacciOptimized Function for calculation fibonacci numbers with cache
func FibonacciOptimized(n int, buffer map[int]int) int {
	if v, ok := buffer[n]; ok {
		return v
	}
	buffer[n] = FibonacciOptimized(n-1, buffer) + FibonacciOptimized(n-2, buffer)
	return buffer[n]
}

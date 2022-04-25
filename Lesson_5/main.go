// Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе
package main

import (
	"fmt"
)

type fibfunc func(int64) int64

//Рекурсивная функция вычисления
func fibRecur(n int64) int64 {
	if n < 2 {
		return 1
	}
	return fibRecur(n-2) + fibRecur(n-1)
}

//Функция формирует слайс фибонначи
func printFib(fib fibfunc, a, b int64) []int64 {
	var fiboSlice []int64
	for i := a; i < b; i++ {
		fiboSlice = append(fiboSlice, fib(i))

	}
	return fiboSlice
}

func main() {
	resultsMaps := make(map[int64][]int64)
	for {
		userInputs, _ := fmt.Println("Введите число: ")
		fmt.Scan(&userInputs)

		if key, ok := resultsMaps[int64(userInputs)]; ok {
			fmt.Printf("Значение уже есть в базе: %d\n", key)
		} else {
			slice := printFib(fibRecur, 0, int64(userInputs))
			resultsMaps[int64(userInputs)] = slice
			fmt.Println(resultsMaps)
		}

	}
}

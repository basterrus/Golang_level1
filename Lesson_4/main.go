package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Вариант 1
//func main() {
//
//	array := []int{3, 4, 1, 2, 5, 7, -1, -5, 0}
//	fmt.Printf("Не отсортированный массив - %d\n", array)
//
//	for i := 1; i < len(array); i++ {
//		x := array[i]
//		j := i
//		for ; j >= 1 && array[j-1] > x; j-- {
//			array[j] = array[j-1]
//		}
//		array[j] = x
//	}
//	fmt.Printf("Отсортированный массив - %d\n", array)
//}

//Вариант 2 ввод чисел через поток

func main() {
	const counts = 5
	inputNums := [counts]int64{}
	fmt.Printf("Введите %d чисел, после ввода каждого числа нажмите Enter: \n", counts)
	scanner := bufio.NewScanner(os.Stdin)
	for i := len(inputNums) - 1; i >= 0; i-- {
		if scanner.Scan() {
			num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			inputNums[i] = num
		}
	}
	fmt.Printf("Исходный массив:%d\n", inputNums)
	for i := 1; i < len(inputNums); i++ {
		x := inputNums[i]
		j := i
		for ; j >= 1 && inputNums[j-1] > x; j-- {
			inputNums[j] = inputNums[j-1]
		}
		inputNums[j] = x
	}
	fmt.Printf("Отсортированный массив :%d\n", inputNums)
}

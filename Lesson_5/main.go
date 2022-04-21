// Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе

package main

import (
	"fmt"
	"os"
)

func fiboGen(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fiboGen(n-1) + fiboGen(n-2)
}

func main() {
	var number uint
	fiboNums := make(map[uint]uint)

	for {
		fmt.Println("Введите число для которого будет сделан рассчет числа фибоначчи (для выхода введите '9999'): ")
		fmt.Scan(&number)

		if number == 9999 {
			os.Exit(1)

		} else if _, ok := fiboNums[number]; ok {
			fmt.Println("Для числа ", number, " число фибоначчи равно = ", fiboNums[number], "значение было рассчитано и сохранено ранее")

		} else {
			fiboNums[number] = fiboGen(number)
			fmt.Println("Для числа ", number, " число фибоначчи равно = ", fiboGen(number), "значение записано в память")

		}
	}
}

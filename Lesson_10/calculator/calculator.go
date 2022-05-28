package calculator

import (
	"fmt"
	"math"
	"os"
)

func Calculator(numOne, numTwo float64, operation string) float64 {

	switch operation {
	case "+":
		return numOne + numTwo
	case "-":
		return numOne - numTwo
	case "*":
		return numOne * numTwo
	case "/":
		//Проверка деления на 0
		if numTwo == 0 {
			fmt.Println("ZeroDivisionError: На ноль делить нельзя! Программа закрыта!")
			os.Exit(1)
		} else {
			return numOne / numTwo
		}
	case "^":
		return math.Pow(numOne, numTwo)

	case "!":
		return float64(factorial(int(numOne)))
	default:
		fmt.Println("Вы выбрали не корректную операцию!")
		os.Exit(1)
	}

	return 0
}

func factorial(n int) float32 {
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return float32(result)
}

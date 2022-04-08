// 2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
// Площадь круга должна вводиться пользователем с клавиатуры.

package main

import (
	"fmt"
	"math"
)

func main() {

	const pi float64 = 3.14
	var s float64

	fmt.Println("Введите площадь круга: ")
	fmt.Scan(&s)

	r := math.Sqrt(s / pi)
	d := r * 2
	l := 2 * pi * r

	fmt.Println("Диаметр круга:", d)
	fmt.Println("Длина окружности:", l)
}

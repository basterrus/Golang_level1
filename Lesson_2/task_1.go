// 1. Напишите программу для вычисления площади прямоугольника.
// Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.

package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Для вычисления площади прямоугольника введите, через пробел,  значение длин его сторон:")
	fmt.Scan(&a, &b)
	fmt.Printf("Площадь прямоугольника равна: %d\n", a*b)

}

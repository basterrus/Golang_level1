// 3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен,
// десятков и единиц в этом числе.

package main

import "fmt"

func main() {
	var num int

	fmt.Println("Введите 3х значное число: ")
	fmt.Scan(&num)

	hundreds := num / 100
	dozens := num / 10 % 10
	units := num % 10

	fmt.Printf("В числе %d\n", num)
	fmt.Printf("Сотен: %d\n", hundreds)
	fmt.Printf("Дестяков: %d\n", dozens)
	fmt.Printf("Единиц: %d\n", units)
}

package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var number int64 = 100
	var numberOfBite int64 = 5
	switchToOne := false
	switch switchToOne {
	case true:
		number |= 1 << numberOfBite
	default:
		number &^= 1 << numberOfBite
	}
	fmt.Println("Ответ: ", number)
}

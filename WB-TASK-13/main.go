package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	first, second := 1, 2
	first, second = second, first
	fmt.Printf("first - %d, second - %d\n", first, second)
	first = first + second
	second = first - second
	first = first - second
	fmt.Printf("first - %d, second - %d", first, second)
}

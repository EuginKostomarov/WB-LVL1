package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

func main() {
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float32)
	for _, val := range temperature {
		key := toRes(val)
		result[key] = append(result[key], val)
	}
	fmt.Println(result)
}

func toRes(val float32) int {
	return int(val) / 10 * 10
}

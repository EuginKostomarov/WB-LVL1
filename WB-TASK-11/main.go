package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	mas1 := []int{1, 7, 3, 5, 3}
	mas2 := []int{3, 6, 1, 9, 3, 7, 2, 4}
	resultMap := make(map[int]bool)
	result := []int{}
	for _, val1 := range mas1 {
		for _, val2 := range mas2 {
			if val1 == val2 {
				resultMap[val1] = true
			}
		}
	}
	for key, value := range resultMap {
		if value {
			result = append(result, key)
		}
	}
	fmt.Println(result)
}

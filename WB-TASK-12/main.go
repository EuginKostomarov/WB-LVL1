package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	set := newSet(slice)
	fmt.Println(set)
}

func newSet(slice []string) []string {
	resMap := make(map[string]bool)
	result := []string{}
	for _, key := range slice {
		resMap[key] = true
	}
	for key, _ := range resMap {
		result = append(result, key)
	}
	return result
}

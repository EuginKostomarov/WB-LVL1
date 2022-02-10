
package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например: 
// abcd — true
// abCdefAaf — false
// 	aabcd — false

func main() {
	str := "cd"
	fmt.Println("result is - ", sleep(str))
}

func sleep(s string) bool {
	str := strings.ToLower(s)
	mapResult := make(map[rune]int)
	for _, symb := range str {
		mapResult[symb]++
	}
	if len(mapResult) != len(str) {
		return false
	}
	return true
}
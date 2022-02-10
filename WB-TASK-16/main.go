package main

import (
	"fmt"
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	mas := []int{342, 234, 2, 0, 23, 4, 5, 7, 9, 2, 35, 63463, 4363, 235}
	sort.Ints(mas)
	fmt.Println(mas)
}

package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10.
//  Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений
var mas [5]int
var summary int

func count(s int, index int) {
	mas[index] = s * s
	summary += s * s
	fmt.Println(summary)
}
func main() {
	mas[0] = 2
	mas[1] = 4
	mas[2] = 6
	mas[3] = 8
	mas[4] = 10
	c := 0
	var wg sync.WaitGroup
	for _, el := range mas {
		wg.Add(1)
		go func() {
			count(el, c)
			c = c + 1
			wg.Done()
		}()
		wg.Wait()
	}

	fmt.Println("Всё готово")

}

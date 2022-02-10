package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.
func count(s int) {
	fmt.Println(s * s)
}
func main() {
	mas := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, el := range mas {
		wg.Add(1)
		go func() {
			count(el)
			wg.Done()
		}()
		wg.Wait()
	}

	fmt.Println("Всё готово")

}

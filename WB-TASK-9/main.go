package main

import "fmt"

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	mas := [...]int{1, 2, 3, 4, 5}

	firstChan := make(chan int)
	secondChan := make(chan int)
	defer close(firstChan)
	defer close(secondChan)

	go toX2(firstChan, secondChan)

	for _, val := range mas {
		firstChan <- val
		fmt.Println(<-secondChan)
	}

}

func toX2(firstChan <-chan int, secondChan chan<- int) {
	for val := range firstChan {
		secondChan <- val * 2
	}
}

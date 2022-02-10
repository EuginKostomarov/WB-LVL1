package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for data := range c {
		time.Sleep(time.Millisecond)
		fmt.Println("Worker is - ", id, ", with message - ", data)
	}
}

func main() {
	c := make(chan int)
	defer close(c)

	//N - количество работяг которые будут генерировать data
	N := 7

	for i := 0; i < N; i++ {
		go worker(i, c)
	}

	for {
		data := rand.Intn(777)
		c <- data
	}
}

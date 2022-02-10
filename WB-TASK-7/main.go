package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.

type chanMap struct {
	key   int
	value int
}

func toData(channel chan chanMap, data map[int]int, sn *sync.Mutex) {
	c := <-channel
	sn.Lock()
	data[c.key] = c.value
	sn.Unlock()
}

func toChanMap(channel chan chanMap) {
	key := rand.Intn(100)
	value := rand.Intn(100)
	channel <- chanMap{key, value}
}

func main() {
	channel := make(chan chanMap)
	defer close(channel)
	data := make(map[int]int)
	sn := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go toChanMap(channel)
		go toData(channel, data, &sn)
	}
	time.Sleep(time.Second)
	fmt.Println(data)
}

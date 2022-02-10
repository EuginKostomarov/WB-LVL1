package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	PrepareFirst()  // 1-ый способ   Проверка по условию
	PrepareSecond() // 2-ой способ  Засыпание по истечению действий
	PrepareThird()  // 3-ий способ   Остановка рутины для ожидания всех остальных
	PrepareFourth() // 4-ый способ  Остановка с использованием контекста выполнения рутины
	PrepareFifth()  // 5-ый способ   Использование флага для условия(не уверен что этот пример нужен, потому что есть первый)
}

func PrepareFirst() {
	quit := make(chan bool)
	defer close(quit)
	go First(quit)
	time.Sleep(time.Second)
	quit <- true
}

func PrepareSecond() {
	chanel := make(chan int)
	go Second(chanel)

	time.Sleep(time.Second)
	close(chanel)
}

func PrepareThird() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			Third(id)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Всё готово!")
}

func PrepareFourth() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := Fourth()
		if err != nil {
			cancel()
		}
	}()
	time.Sleep(time.Second)
	Fourth1(ctx)
}

func PrepareFifth() {
	flag := false
	go Fifth(&flag)
	time.Sleep(time.Second)
	flag = true
	time.Sleep(time.Second)
}

func First(quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Printf("Не работает\n")
			return
		default:
			fmt.Printf("Работает\n")
		}
	}
}

func Second(chanel chan int) {
	for {
		fmt.Printf("do something with chanel\n")
	}
}

func Third(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Горутина %d, отметилась\n", id)
	}
}

func Fourth() error {
	time.Sleep(time.Second)
	return errors.New("Ошибка: Слишком долго")
}

func Fourth1(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Printf("Last process too long")
	default:
		fmt.Printf("Normal mode")
	}
}

func Fifth(flag *bool) {
	for {
		if *flag {
			fmt.Println("Выполнено")
			return
		}
	}
}

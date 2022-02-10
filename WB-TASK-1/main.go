package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
	university string
}

// TASK1
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s \n", h.name)
}

func main() {
	s := Action{Human{name: "Евгений", age: 23}, "МТУСИ"}
	s.SayHi()
	s.Human.SayHi()

}

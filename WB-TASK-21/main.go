
package main

import "fmt"

type kilo int

// Реализовать паттерн «адаптер» на любом примере.

func main() {
	gramm := 500
	something(toKilo(gramm))
}

func toKilo(data int) kilo {
	return kilo(data * 1000)
}

func something(v kilo) {
	fmt.Printf("value - %d, type - %T", v, v)
}

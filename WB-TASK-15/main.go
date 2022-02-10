package main

import (
	"fmt"
	"unsafe"
)

var justString string

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {

		v += "A"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	res := make([]rune, 0, 100)
	res = []rune(v)[:100]
	justString2 := string(res)

	fmt.Println("v = ", v)
	fmt.Println("justString = ", justString)
	fmt.Println("justString2 = ", justString2)

	fmt.Println("len(v) = ", len(v))
	fmt.Println("len(justString) = ", len(justString))
	fmt.Println("len(justString2) = ", len(justString2))

	fmt.Println("cap(v) = ", cap([]byte(v)))
	fmt.Println("cap(justString) = ", cap([]byte(justString)))
	fmt.Println("cap(justString2) = ", cap([]byte(justString2)))

	fmt.Println("sizeof(v) = ", unsafe.Sizeof(v))
	fmt.Println("sizeof(justString) = ", unsafe.Sizeof(justString))
	fmt.Println("sizeof(justString2) = ", unsafe.Sizeof(justString2))

	fmt.Printf("Pointer(v) = %p\n", &v)
	fmt.Printf("Pointer(justString) = %p\n", &justString)
	fmt.Printf("Pointer(justString2) = %p\n", &justString2)
}

func main() {
	someFunc()
}

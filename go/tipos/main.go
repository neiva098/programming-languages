package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("O tipo inteiro e", reflect.TypeOf(7787))

	// uint8 uint16 uint32 uint64 +
	var x uint8 = 255
	fmt.Println(x)

	// int8 int16 int32 int64

	var y rune = 'a'
	fmt.Println(y)

	// float32, float64

	const u = float64(66)

	fmt.Printf("%f", u)

	// bool

	// string
	fmt.Println("\ntamanho", len("ola"))
}

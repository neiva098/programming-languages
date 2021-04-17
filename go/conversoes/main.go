package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	fmt.Println(int(nota))

	// cuidado
	fmt.Println("teste " + string(97))

	fmt.Println("Teste " + strconv.Itoa(124))

	num, _ := strconv.Atoi("123")

	fmt.Println("Teste", num)

	b, _ := strconv.ParseBool("true") // 1 == verdadeiro

	if b {
		fmt.Println("passou")
	}
}

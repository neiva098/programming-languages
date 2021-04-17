package main

import "fmt"

func main() {
	var a byte = 3

	fmt.Println(a)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2

	x, y := 2, 1
	x, y = y, x

	fmt.Println(x, y)
}

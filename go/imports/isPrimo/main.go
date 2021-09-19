package main

import (
	"fmt"
	math "isPrimo/math"
	"os"
	"strconv"
)

func main() {
	number, _ := strconv.Atoi(os.Args[1])

	fmt.Println(math.IsPrimo(number))
}

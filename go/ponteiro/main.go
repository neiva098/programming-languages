package main

import "fmt"

// sem aritmetica com ponteiros

func main() {
	i := 1

	var p *int = nil

	p = &i
	*p++

	fmt.Println(i, &i, *p, p, &p)
}

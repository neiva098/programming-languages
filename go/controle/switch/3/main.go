package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float64, float32:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "nao sei"
	}
}

func main() {
	fmt.Println(tipo(2.44))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}

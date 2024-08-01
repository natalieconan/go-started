package main

import (
	"fmt"
)

type Number interface {
	int | int8 | float32 | float64
}

func add[T Number](a ...T) T {
	var sum T = 0
	for _, v := range a {
		sum += v
	}

	return sum
}

func cusPrint[T any](s ...T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	str := "Hello World"
	str2 := "Something"

	cusPrint(str, str2)

	fmt.Println(add(3, 4, 5, 6, 7))
}

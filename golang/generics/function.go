package main

import "fmt"

// function with generic type
func print[T any](s []T) {
	for _, v := range s {
		println(v)
	}
}

func double[T int | float32](s T) T {
	return s * 2
}

func main() {
	print([]int{1, 2, 3})
	print([]string{"a", "b", "c"})
	v := double[int](5)
	fmt.Println(v)
	w := double[float32](5.6)
	fmt.Println(w)
}

package main

import "fmt"

// for params can declare the type individaully
func add(a int, b int) int {
	return a + b
}

// or you can just declare the type at the end of the same type params
func sub(a, b int) int {
	return a - b
}

func main() {
	res := add(1, 2)
	res2 := sub(1, 2)

	fmt.Println("res:", res)
	fmt.Println("res2:", res2)
}

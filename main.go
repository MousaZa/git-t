package main

import "fmt"

func main() {
	fmt.Printf("1 + 2 = %d\n", add(1, 2))
	fmt.Printf("5 - 3 = %d\n", subtract(5, 3))
	fmt.Printf("4 * 2 = %d\n", multiply(4, 2))
	fmt.Printf("8 / 2 = %d\n", divide(8, 2))
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func solve(a, b, c int) int {
	return add(multiply(a, b), c)
}

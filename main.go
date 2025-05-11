package main

import "fmt"

func main() {
	fmt.Printf("1 + 2 = %d\n", add(1, 2))
	fmt.Printf("5 - 3 = %d\n", subtract(5, 3))
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

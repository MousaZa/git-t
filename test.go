package main

func testAdd() {
	result := add(1, 2)
	if result != 3 {
		panic("add(1, 2) = " + string(result) + ", want 3")
	}
}

func testSubtract() {
	result := subtract(5, 3)
	if result != 2 {
		panic("subtract(5, 3) = " + string(result) + ", want 2")
	}
}

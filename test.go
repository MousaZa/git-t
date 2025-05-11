package main

func testAdd() {
	result := add(1, 2)
	if result != 3 {
		panic("add(1, 2) = " + string(result) + ", want 3")
	}
}

package main

import "fmt"

/*

func name(argument) return_type{
	body
}

*/

func main() {
	res := add(3, 5)

	fmt.Println(res)
	fmt.Println(proAdder(3, 8, 8, 8, 3, 4, 5, 6))
}

// pro functions
func proAdder(values ...int) (int, string) {
	total := 0
	for i := range values {
		 total += i
	}
	return total, "Solved"
}

func add(a int, b int) int {
	return a+ b
}
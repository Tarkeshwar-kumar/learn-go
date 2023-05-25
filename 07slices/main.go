package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice = []int{2, 4, 8, 0, 9}

	sort.Ints(slice)
	fmt.Println(slice)


	slice2 := make([]int, 3)
	slice2[0] = 9; slice2[1] = 5; slice2[2] = 0
	slice2 = append(slice, 4, 2)
	sort.Ints(slice2)
	fmt.Println(slice2)

	// to remove index 3 from slice

	var index int = 3

	slice2 = append(slice2[:index], slice2[index+ 1:]...)
	fmt.Println(slice2)
}
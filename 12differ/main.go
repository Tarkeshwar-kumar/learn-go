package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")

	defer deff()
}

func deff(){
	for i := 0; i<5; i++ {
		defer fmt.Println(i)
	}
}
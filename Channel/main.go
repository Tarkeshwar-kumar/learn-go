package main

import "fmt"

func main() {
	var first_channel chan int
	fmt.Println("Value of the channel is ", first_channel)
	fmt.Printf("Type of the channel is %T\n", first_channel)

	second_channel := make(chan int)
	fmt.Println("Value of the channel is ", second_channel)
	fmt.Printf("Type of the channel is %T", second_channel)
}
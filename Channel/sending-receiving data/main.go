package main

import "fmt"

func main() {
	fmt.Println("Strating the main function")
	second_channel := make(chan int)

	go myFunction(second_channel)
	second_channel <- 32
	
	fmt.Println("Finishing the main function execution")
}

func myFunction(channel chan int) {
	fmt.Println(<-channel)
}
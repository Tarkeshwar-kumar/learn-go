package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Enter a number")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')  // this is comma ok syntax
	fmt.Println("Enterd input is ", input)
	fmt.Println("Error occured during reading input is ", err)
}
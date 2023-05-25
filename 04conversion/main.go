package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Entered number is ", input)
	number, err_on_converting := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err_on_converting !=  nil {
		fmt.Println("error is", err_on_converting)
	} else {
		fmt.Println("converted number is ", number)
		fmt.Printf("%T", number)
	}
}
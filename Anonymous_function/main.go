package main

import (
	"fmt"
)

func main() {
	//simple anonymous function
	func () {
		fmt.Println("BAKKAAAAAA.....")
	}()

	// storing variable and calling it
	variable := func(){
		fmt.Println("Hello World")
	}
	variable()
	// type of variable is of func type
	fmt.Printf("Type of variable1 is %T\n", variable)

	//passing argument to the function
	variable2 := func(say_hello string){
		fmt.Println("Hello World")
	}
	variable2("Hello World!")

	// passing anonymous function as an argument into another function
	variable3 := func (p, q string) string {
		return p+ q+ "Not Welcome"
	}
	test_function(variable3)

	// returning an anonymous function from another function
	variable4 := Returning_function()
	fmt.Printf(variable4("Simon ", "comission "))
}

func test_function(a_func func(p, q string) string){
	fmt.Println(a_func ("Saloo ", "Bhai "))
}

func Returning_function() func (p, q string) string{
	to_return := func (p, q string) string {
		return p+ q+ "go back"
	}
	return to_return
}
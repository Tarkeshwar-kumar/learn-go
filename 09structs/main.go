package main

import "fmt"

type User struct {
	Name       string
	age        string
	Profession string
}

func main() {
	// no inheritance in go, no super and parent
	Chelsi := User{"Chelsi", "22", "Doctor"}
	fmt.Println(Chelsi)
	fmt.Println(Chelsi.age)
}

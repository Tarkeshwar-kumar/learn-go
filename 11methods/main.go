package main

import "fmt"

type User struct {
	Name       string
	age        int
	Profession string
	isOnline   bool
}

func main() {
	Ashok := User{"Ashok", 22, "Go Developer", true}

	fmt.Println(Ashok.age)
	fmt.Println(Ashok.Name)
	fmt.Println(Ashok.Profession)
	fmt.Println(Ashok.isOnline)

	Ashok.getUser()
}


func (u User) getUser(){
	fmt.Println("Name of the user is : ", u.Name)
	fmt.Println("Age of the user is : ", u.age)
	fmt.Println("Profession of the user is : ", u.Profession)
	if u.isOnline {
		fmt.Println("User is online")
	} else{
		fmt.Println("User is offline")
	}
}
package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name       string `json:"employeename"`
	Id         string `json:"employeeid"`
	Age        int
	Password   string `json:"-"`
	Department string `json:"department,omitempty"`
	Skills     []string `json:"skills"`
}

func main() {
	employee := []Employee{
		{"Chelsi", "12ds32", 22, "abc@123", "Medical", []string{"surgery", "prescription", "diagnosis"}},
		{"Tarkeshwar", "da12@32", 22, "davidMorin", "Analytics", []string{"Python", "SQL", "PowerBI"}},
		{"Tom", "2334reer", 44, "fdefe322", "Developer", []string{"Go", "AWS", "Docker"}},
	}
	JosnEmployee, err := json.MarshalIndent(employee, "", "\t")

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println(string(JosnEmployee))
}
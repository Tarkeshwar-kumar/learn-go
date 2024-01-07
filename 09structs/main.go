package main

import "fmt"

type Profession struct {
	profession_name string
	skill []string
	field string
}

type Person struct {
	name       string
	age        int
	profession Profession
}

func (profession *Profession) describe_profession_skill() {
	fmt.Printf("profession %s required following skillset:\n", profession.profession_name)
	for _, skill := range profession.skill {
		fmt.Println("-> ", skill)
	}
}

func (person *Person) introduce() {
	fmt.Printf("Hello World!! My name is %s and my age is %d\n", person.name, person.age)
	fmt.Printf("I am %s, I work in %s field\n", person.profession.profession_name, person.profession.field)
	person.profession.describe_profession_skill()
}

func (profession *Profession) change_skill(index int, skill string) {
	profession.skill[index] = skill
}

func main() {
	// no inheritance in go, no super and parent
	go_dev_skill := []string {"Go", "Git", "gRPC", "gin", "Docker", "K8s"}
	prof := Profession{"Go developer", go_dev_skill, "software Engineer"}
	person := Person{"Tarak", 22, prof}
	fmt.Println(person)
	person.introduce()
	prof.change_skill(5, "Kubernetes")
	person.introduce()
}	
	

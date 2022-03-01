package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Talk() {
	fmt.Println("Kono", p.Name, "Da!")
}

type StandUser struct {
	Person
	StandName string
}

func main() {
	d := StandUser{
		Person: Person{
			Name: "Dio",
		},
		StandName: "The World",
	}

	d.Talk()
	fmt.Println(d.Name, "Stand name:", d.StandName)
}

package main

import "fmt"

type Person interface {
	Walk()
	Run()
}

type Superman struct {
	Name string
}

func (s Superman) Walk() {
	fmt.Println(s.Name, "(superman) is walking.")
}

func (s Superman) Run() {
	fmt.Println(s.Name, "(superman) is running.")
}

func DoSomething(p Person) {
	p.Walk()
	// p.Run()
}

func main() {
	s := Superman{
		Name: "Clark Kent",
	}

	DoSomething(s)
}

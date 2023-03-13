package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Greet() {
	fmt.Printf("Hello, my name is %s\n", h.Name)
}

type Action struct {
	Human
	Action string
}

func main() {
	a := Action{
		Human: Human{
			Name: "John",
			Age:  30,
		},
		Action: "running",
	}
	a.Greet()
}

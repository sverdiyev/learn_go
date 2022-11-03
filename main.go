package main

import "fmt"

type Animal interface {
	Says() string
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name  string
	Color string
}

//Interface is implemented because the type implements its methods
//Not because the type implements the interface

// Best practice - pass the reference

func (d *Dog) Says() string {
	return "My name is " + d.Name
}
func (c *Cat) Says() string {
	return "My name is " + c.Name
}

func introduce(animal Animal) {
	fmt.Println(animal.Says())
}

func main() {
	myDog := Dog{"DOG", "Boiboi"}
	notMyCat := Cat{"CAT", "girgirl"}
	introduce(&myDog)
	introduce(&notMyCat)
}

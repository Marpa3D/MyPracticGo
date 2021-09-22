package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main()  {
	dog := Dog{
		Name: "Bublik",
		Breed: "Buldog",
	}

	PrintInfo(dog)
}

func (d Dog) Says() string {
	return "Gaaf!)"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal)  {
	log.Println("Это животное говорит:", a.Says(), "И у него", a.NumberOfLegs(), "ноги.")
}
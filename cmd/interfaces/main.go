// Практика создания интерфейсов в языке Go от Вячеслава Шарова

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

	gorilla := Gorilla{
		Name: "Bob",
		Color: "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(dog)
	PrintInfo(gorilla)
}

// Реализация для типа Dog
func (d Dog) Says() string {
	return "Гааф!)"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

// Реализация для типа Gorilla
func (g Gorilla) Says() string {
	return "Я большая Горилла!)))"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal)  {
	log.Println("Это животное говорит:", a.Says(), "И у него", a.NumberOfLegs(), "ноги.")
}
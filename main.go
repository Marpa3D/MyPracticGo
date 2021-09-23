package main

import (
	"MyPracticGo/cmd/packages"
	"log"
)

func main()  {
	var myVar packages.SomeThing
	myVar.Name = "Slava"

	log.Println(myVar.Name)
}

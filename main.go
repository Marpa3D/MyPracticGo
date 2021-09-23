package main

import (
	"MyPracticGo/cmd/channels"
	"log"
)

const numPool = 1024

func CalcuateValues(inChan chan int)  {
	randomNumber := channels.RandomNumbers(numPool)
	inChan <- randomNumber
}

func main()  {

	inChan := make(chan int)
	defer close(inChan)

	go CalcuateValues(inChan)
	num := <- inChan

	log.Println(num)
}

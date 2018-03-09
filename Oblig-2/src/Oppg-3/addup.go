package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"strconv"
)

func main() {

	sigInt := make(chan os.Signal, 1)
	signal.Notify(sigInt, os.Interrupt)

	go func() {
		<-sigInt
		fmt.Println("Interruption signal recived")
		os.Exit(1)
	}()



	channel := make(chan int)
	go readInput(channel)
	time.Sleep(5 * time.Second)
	go addUp(channel)
	time.Sleep(5 * time.Second)
}

func readInput(channel chan int) {

	var number1 string
	var number2 string

	fmt.Println("Enter number: ")
	fmt.Scan(&number1)
	intNumber1, err := strconv.Atoi(number1)
	if err != nil{
		fmt.Println("Please input a number. Shutting down.")
		os.Exit(0)
	}

	fmt.Println("Enter number to add to the first number: ")
	fmt.Scan(&number2)
	intNumber2, err := strconv.Atoi(number2)
	if err != nil{
		fmt.Println("Please input a number. Shutting down.")
		os.Exit(0)
	}

	channel <- intNumber1
	channel <- intNumber2

	result := <-channel
	fmt.Println("Result:",number1,"+",number2,"=", result)

}

func addUp(channel chan int) {

	number1, number2 := <-channel, <-channel

	result := (number1 + number2)

	channel <- result
}

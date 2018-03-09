package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	sigInt := make(chan os.Signal, 1)
	signal.Notify(sigInt, os.Interrupt)

	go func() {
		<-sigInt
		fmt.Println("Interruption signal recived, terminating program...... ")
		time.sleep(2*time.Second)
		fmt.Println("Terminated")
		os.Exit(1)
	}()

	channel := make(chan int)
	go readInput(channel)
	time.Sleep(5 * time.Second)
	go addUp(channel)
	time.Sleep(5 * time.Second)
}

func readInput(channel chan int) {

	var number1 int
	var number2 int

	fmt.Println("Enter number: ")
	fmt.Scan(&number1)
	fmt.Println("Enter number to add to the first number: ")
	fmt.Scan(&number2)

	channel <- number1
	channel <- number2

	result := <-channel
	fmt.Println("Result:",number1,"+",number2,"=", result)

}

func addUp(channel chan int) {

	number1, number2 := <-channel, <-channel

	result := (number1 + number2)

	channel <- result
}

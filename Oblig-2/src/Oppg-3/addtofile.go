package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()


	var number1 int

	var number2 int



	fmt.Println("Enter number: ")

	fmt.Scan(&number1)

	fmt.Println("Enter number to add to the first number: ")

	fmt.Scan(&number2)

	t := strconv.Itoa(number1)
	n := strconv.Itoa(number2)

	fmt.Fprintf(file, t)
	fmt.Fprintf(file, n)
}

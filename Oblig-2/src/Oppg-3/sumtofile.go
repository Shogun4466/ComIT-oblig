package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

func sumFromFile()  {
	line, err := readLines("result.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}


	input1 := line[0]
	input2 := line[1]

	number1,_ := strconv.Atoi(input1)
	number2,_ := strconv.Atoi(input2)

	result := number1 + number2


	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	if _, err := fmt.Fprintf(file,"\n%d\n", result); err != nil {
		log.Fatal("Failed to write to file", err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

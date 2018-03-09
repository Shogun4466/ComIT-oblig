package main
import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os/signal"
	"time"
)

func main() {

	sigInt := make(chan os.Signal, 1)
	signal.Notify(sigInt, os.Interrupt)

	go func() {
		<-sigInt
		fmt.Println("Interruption signal recived, terminating program...... ")
		time.Sleep(1*time.Second)
		fmt.Println("Terminated")
		os.Exit(1)
	}()

	writeToFile()
	sumFromFile()
	readResult("result.txt")
}

func writeToFile() {

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

	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Failed to create file", err)
	}
	defer file.Close()

	f, err := os.OpenFile("result.txt",os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	if _, err := fmt.Fprintf(f, "%d\n%d", intNumber1, intNumber2); err != nil {
		log.Fatal("Failed to write to file", err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func readResult(path string) {
	data, err := ioutil.ReadFile(path)
	checkErr(err)

	tempData := string(data)
	stringData := strings.Split(tempData, "\n")
	tempNumber1 := stringData[len(stringData)-4]
	tempNumber2 := stringData[len(stringData)-3]
	tempResult := stringData[len(stringData)-2]

	number1, err := strconv.Atoi(tempNumber1)
	number2, err := strconv.Atoi(tempNumber2)
	result, err := strconv.Atoi(tempResult)

	checkErr(err)

	fmt.Println("Result from file:",number1,"+",number2,"=", result)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

package main



import (
	"os"

	"strconv"

	"fmt"

	"log"
)



func main() {

	input1 := os.Args[1]

	input2 := os.Args[2]

	tall1, err := strconv.Atoi(input1)
	tall2, err := strconv.Atoi(input2)

	if err != nil {

		log.Fatal(err)

	}

	s := make(chan int)

	go noe(tall1, tall2, s)

	sum := <-s

	fmt.Println(sum)
}

func noe(tall1 int, tall2 int, s chan int) {

	sum := tall1 + tall2

	s <- sum
}

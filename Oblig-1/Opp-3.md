package main
import (
	"fmt"

	"os"

	"os/signal"

	"syscall"
)
func main() {
	indefLoop()
}
func indefLoop() {
	c := make(chan os.Signal, 1)

	signal.Notify(c,

		syscall.SIGINT,)

	for {

		fmt.Println("Loooooooooooooop")

		a := <-c

		switch a {

		case syscall.SIGINT:

			fmt.Println("You shall not pass SIGINT")
		}
		break
	}
}

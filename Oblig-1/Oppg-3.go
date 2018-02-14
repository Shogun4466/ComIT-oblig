//I denne koden så opretes det en uendelig loop som blir avslutet når datamaskinen får inputt av type CTRL+fn+f2. Når inputet blir tatt imot,
//så skriver koden meldingen "You shall not pass SIGINT
// Når programet Kjører, så skrives det ut en melding "Loooooooop" intil programet stopes. Programet tar ikke opp noe av prossesorkraften 
// under kjøringen, men det tar 0,9MB. 
package main 
//Bruker import for å inhente alle nødvendige funksjoner
import ( 
"fmt"
	
"os"

"os/signal"

"syscall"
) 
func main() { 
	indefLoop() 
} 
//Her settes signalet opp
func indefLoop() { 

c := make(chan os.Signal, 1)
signal.Notify(c,

	syscall.SIGINT,)
//En uendlig for loop som skriver "Loooooooooop" og sjekker etter syscall
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

package main
import (
	"fmt"
	"net"
	"bufio"
)
//For å teste serverfunksjonalitet kan udp og tcp skrives inn avhengig av hvilken server du skal koble til.
func main() {
	p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:17")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi Server, How you doin'?")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}

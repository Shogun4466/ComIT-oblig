package main

import (
	"fmt"
	"net"
	"log"
	"time"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_,err := conn.WriteToUDP([]byte("Quote of the Day - UDP port 17"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	go tcpServer ()
	go udpServer ()
	time.Sleep(time.Second * 100000)
}

func udpServer() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 17,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}

func tcpServer(){
	l, err := net.Listen("tcp", "127.0.0.1:17")
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go func(c net.Conn) {
				c.Write([]byte("Quote of the day - TCP port 17"))
				c.Close()
			}(conn)
		}
	}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprintln(conn, "how is your day?")
		fmt.Fprintf(conn, "%v", "Good, I hope\n")

		conn.Close()
	}
}
